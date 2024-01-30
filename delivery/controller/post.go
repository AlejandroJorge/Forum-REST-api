package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/AlejandroJorge/forum-rest-api/delivery"
	"github.com/AlejandroJorge/forum-rest-api/domain"
	"github.com/AlejandroJorge/forum-rest-api/util"
	"github.com/gorilla/mux"
)

type postController struct {
	serv domain.PostService
}

func NewPostController(serv domain.PostService) postController {
	return postController{serv: serv}
}

func (con postController) GetByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idStr, ok := params["id"]
	if !ok {
		delivery.WriteResponse(w, http.StatusBadRequest, "No provided ID")
		return
	}

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		delivery.WriteResponse(w, http.StatusBadRequest, "ID provided isn't a number")
		return
	}

	post, err := con.serv.GetByID(uint(id))
	if err == util.ErrEmptySelection {
		delivery.WriteResponse(w, http.StatusNotFound, "There's no post with this ID")
		return
	}
	if err != nil {
		delivery.WriteResponse(w, http.StatusInternalServerError, "Couldn't retrieve post")
	}

	delivery.WriteJSONResponse(w, http.StatusOK, post)
}

func (con postController) GetByUserID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idStr, ok := params["id"]
	if !ok {
		delivery.WriteResponse(w, http.StatusBadRequest, "No provided ID")
		return
	}

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		delivery.WriteResponse(w, http.StatusBadRequest, "ID provided isn't a number")
		return
	}

	post, err := con.serv.GetByUser(uint(id))
	if err == util.ErrEmptySelection {
		delivery.WriteResponse(w, http.StatusNotFound, "There's no post with this ID")
		return
	}
	if err != nil {
		delivery.WriteResponse(w, http.StatusInternalServerError, "Couldn't retrieve post")
	}

	delivery.WriteJSONResponse(w, http.StatusOK, post)
}

func (con postController) GetPopular(w http.ResponseWriter, r *http.Request) {
	var getPopularReq struct {
		Interval string `json:"Interval"`
	}
	err := delivery.ReadJSONRequest(r, &getPopularReq)
	if err != nil {
		delivery.WriteJSONResponse(w, http.StatusBadRequest, "Incorrect request format")
		return
	}

	var getPopular func() ([]domain.Post, error)
	switch getPopularReq.Interval {
	case "Ever":
		getPopular = con.serv.GetPopularAllTime
		break
	case "Month":
		getPopular = con.serv.GetPopularLastMonth
		break
	case "Week":
		getPopular = con.serv.GetPopularLastWeek
		break
	case "Day":
		getPopular = con.serv.GetPopularToday
		break
	default:
		delivery.WriteResponse(w, http.StatusBadRequest, fmt.Sprintf("Incorrect option: %s", getPopularReq.Interval))
		return
	}

	posts, err := getPopular()
	if err == util.ErrEmptySelection {
		delivery.WriteResponse(w, http.StatusNotFound, "No posts matched criteria")
		return
	}
	if err != nil {
		delivery.WriteResponse(w, http.StatusInternalServerError, "Couldn't retrieve posts")
	}

	delivery.WriteJSONResponse(w, http.StatusOK, posts)
}

func (con postController) Create(w http.ResponseWriter, r *http.Request) {
	var createReq struct {
		OwnerID     uint   `json:"OwnerID"`
		Title       string `json:"Title"`
		Description string `json:"Description"`
		Content     string `json:"Content"`
	}
	err := delivery.ReadJSONRequest(r, &createReq)
	if err != nil {
		delivery.WriteResponse(w, http.StatusBadRequest, "Incorrect request format")
		return
	}

	id, err := con.serv.CreateNew(struct {
		OwnerID     uint
		Title       string
		Description string
		Content     string
	}{
		OwnerID:     createReq.OwnerID,
		Title:       createReq.Title,
		Description: createReq.Description,
		Content:     createReq.Content,
	})
	if err == util.ErrNoCorrespondingUser {
		delivery.WriteResponse(w, http.StatusBadRequest, "Owner doesn't exist")
		return
	}
	if err != nil {
		delivery.WriteResponse(w, http.StatusInternalServerError, "Couldn't create post")
		return
	}

	response := struct {
		ID uint `json:"ID"`
	}{
		ID: id,
	}
	delivery.WriteJSONResponse(w, http.StatusOK, response)
}

func (con postController) Update(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idStr, ok := params["id"]
	if !ok {
		delivery.WriteResponse(w, http.StatusBadRequest, "No provided ID")
		return
	}

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		delivery.WriteResponse(w, http.StatusBadRequest, "ID provided isn't a number")
		return
	}

	var updateReq struct {
		UpdatedTitle       string `json:"Title"`
		UpdatedDescription string `json:"Description"`
		UpdatedContent     string `json:"Content"`
	}
	err = delivery.ReadJSONRequest(r, &updateReq)
	if err != nil {
		delivery.WriteResponse(w, http.StatusBadRequest, "Incorrect request format")
		return
	}

	err = con.serv.Update(uint(id), struct {
		UpdatedTitle       string
		UpdatedDescription string
		UpdatedContent     string
	}{
		UpdatedTitle:       updateReq.UpdatedTitle,
		UpdatedDescription: updateReq.UpdatedDescription,
		UpdatedContent:     updateReq.UpdatedContent,
	})
	if err != nil {
		delivery.WriteResponse(w, http.StatusInternalServerError, "Couldn't update")
		return
	}

	delivery.WriteResponse(w, http.StatusOK, "Updated post successfully")
}

func (con postController) AddLike(w http.ResponseWriter, r *http.Request) {
	var addLikeReq struct {
		UserID uint `json:"UserID"`
		PostID uint `json:"PostID"`
	}
	err := delivery.ReadJSONRequest(r, &addLikeReq)
	if err != nil {
		delivery.WriteResponse(w, http.StatusBadRequest, "Incorrect request format")
		return
	}

	err = con.serv.AddLike(addLikeReq.UserID, addLikeReq.PostID)
	if err != nil {
		delivery.WriteResponse(w, http.StatusInternalServerError, "Couldn't create like")
		return
	}

	delivery.WriteResponse(w, http.StatusOK, "Like created successfully")
}

func (con postController) DeleteLike(w http.ResponseWriter, r *http.Request) {
	var deleteLikeReq struct {
		UserID uint `json:"UserID"`
		PostID uint `json:"PostID"`
	}
	err := delivery.ReadJSONRequest(r, &deleteLikeReq)
	if err != nil {
		delivery.WriteResponse(w, http.StatusBadRequest, "Incorrect request format")
		return
	}

	err = con.serv.DeleteLike(deleteLikeReq.UserID, deleteLikeReq.PostID)
	if err != nil {
		delivery.WriteResponse(w, http.StatusInternalServerError, "Couldn't delete like")
		return
	}

	delivery.WriteResponse(w, http.StatusOK, "Like deleted successfully")
}

func (con postController) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idStr, ok := params["id"]
	if !ok {
		delivery.WriteResponse(w, http.StatusBadRequest, "No provided ID")
		return
	}

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		delivery.WriteResponse(w, http.StatusBadRequest, "ID provided isn't a number")
		return
	}

	err = con.serv.Delete(uint(id))
	if err != nil {
		delivery.WriteResponse(w, http.StatusInternalServerError, "Couldn't delete post")
		return
	}

	delivery.WriteResponse(w, http.StatusOK, "Post deleted successfully")
}

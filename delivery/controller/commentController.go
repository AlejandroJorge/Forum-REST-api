package controller

import (
	"net/http"
	"strconv"

	"github.com/AlejandroJorge/forum-rest-api/delivery"
	"github.com/AlejandroJorge/forum-rest-api/domain"
	"github.com/AlejandroJorge/forum-rest-api/util"
	"github.com/gorilla/mux"
)

type commentController struct {
	serv domain.CommentService
}

func NewCommentController(serv domain.CommentService) commentController {
	return commentController{serv: serv}
}

func (con commentController) Create(w http.ResponseWriter, r *http.Request) {
	var createReq struct {
		UserID  uint   `json:"UserID"`
		PostID  uint   `json:"PostID"`
		Content string `json:"Content"`
	}
	err := delivery.ReadJSONRequest(r, &createReq)
	if err != nil {
		delivery.WriteJSONResponse(w, http.StatusInternalServerError, "Couldn't create comment")
		return
	}

	id, err := con.serv.CreateNew(struct {
		UserID  uint
		PostID  uint
		Content string
	}{
		UserID:  createReq.UserID,
		PostID:  createReq.PostID,
		Content: createReq.Content,
	})
	if err != nil {
		delivery.WriteResponse(w, http.StatusInternalServerError, "Couldn't create comment")
		return
	}

	response := struct {
		ID uint `json:"ID"`
	}{
		ID: id,
	}
	delivery.WriteJSONResponse(w, http.StatusCreated, response)
}

func (con commentController) GetByID(w http.ResponseWriter, r *http.Request) {
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

	comment, err := con.serv.GetByID(uint(id))
	if err == util.ErrEmptySelection {
		delivery.WriteResponse(w, http.StatusNotFound, "There's no comment with this ID")
		return
	}
	if err != nil {
		delivery.WriteResponse(w, http.StatusInternalServerError, "Couldn't retrieve comment")
		return
	}

	delivery.WriteJSONResponse(w, http.StatusOK, comment)
}

func (con commentController) GetByUser(w http.ResponseWriter, r *http.Request) {
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

	posts, err := con.serv.GetByUser(uint(id))
	if err == util.ErrEmptySelection {
		delivery.WriteResponse(w, http.StatusNotFound, "There are no comments for this user")
		return
	}
	if err != nil {
		delivery.WriteResponse(w, http.StatusInternalServerError, "Couldn't retrieve comments")
		return
	}

	delivery.WriteJSONResponse(w, http.StatusOK, posts)
}

func (con commentController) GetByPost(w http.ResponseWriter, r *http.Request) {
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

	posts, err := con.serv.GetByPost(uint(id))
	if err == util.ErrEmptySelection {
		delivery.WriteResponse(w, http.StatusNotFound, "There are no comments for this post")
		return
	}
	if err != nil {
		delivery.WriteResponse(w, http.StatusInternalServerError, "Couldn't retrieve comments")
		return
	}

	delivery.WriteJSONResponse(w, http.StatusOK, posts)
}

func (con commentController) Update(w http.ResponseWriter, r *http.Request) {
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
		Content string `json:"Content"`
	}
	err = delivery.ReadJSONRequest(r, &updateReq)
	if err != nil {
		delivery.WriteResponse(w, http.StatusBadRequest, "Incorrect request format")
		return
	}

	con.serv.Update(uint(id), updateReq.Content)
}

func (con commentController) CreateLike(w http.ResponseWriter, r *http.Request) {
	var createLikeReq struct {
		UserID    uint `json:"UserID"`
		CommentID uint `json:"CommentID"`
	}
	err := delivery.ReadJSONRequest(r, &createLikeReq)
	if err != nil {
		delivery.WriteResponse(w, http.StatusBadRequest, "Incorrect request format")
		return
	}

	err = con.serv.AddLike(createLikeReq.UserID, createLikeReq.CommentID)
	if err != nil {
		delivery.WriteResponse(w, http.StatusInternalServerError, "Couldn't create like")
		return
	}

	delivery.WriteResponse(w, http.StatusOK, "Like created successfully")
}

func (con commentController) DeleteLike(w http.ResponseWriter, r *http.Request) {
	var deleteLikeReq struct {
		UserID    uint `json:"UserID"`
		CommentID uint `json:"CommentID"`
	}
	err := delivery.ReadJSONRequest(r, &deleteLikeReq)
	if err != nil {
		delivery.WriteResponse(w, http.StatusBadRequest, "Incorrect request format")
		return
	}

	err = con.serv.DeleteLike(deleteLikeReq.UserID, deleteLikeReq.CommentID)
	if err != nil {
		delivery.WriteResponse(w, http.StatusInternalServerError, "Couldn't delete like")
		return
	}

	delivery.WriteResponse(w, http.StatusOK, "Like deleted successfully")
}

func (con commentController) Delete(w http.ResponseWriter, r *http.Request) {
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
		delivery.WriteResponse(w, http.StatusInternalServerError, "Couldn't delete comment")
		return
	}

	delivery.WriteResponse(w, http.StatusOK, "Comment deleted successfully")
}

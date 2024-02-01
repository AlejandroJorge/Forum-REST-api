package controller

import (
	"net/http"

	"github.com/AlejandroJorge/forum-rest-api/delivery"
	"github.com/AlejandroJorge/forum-rest-api/domain"
	"github.com/AlejandroJorge/forum-rest-api/service"
)

type PostController interface {
	Create(w http.ResponseWriter, r *http.Request)

	Delete(w http.ResponseWriter, r *http.Request)

	UpdateTitle(w http.ResponseWriter, r *http.Request)

	UpdateDescription(w http.ResponseWriter, r *http.Request)

	UpdateContent(w http.ResponseWriter, r *http.Request)

	GetByID(w http.ResponseWriter, r *http.Request)

	GetByUser(w http.ResponseWriter, r *http.Request)

	GetPopularToday(w http.ResponseWriter, r *http.Request)

	GetPopularLastWeek(w http.ResponseWriter, r *http.Request)

	GetPopularLastMonth(w http.ResponseWriter, r *http.Request)

	GetPopularAllTime(w http.ResponseWriter, r *http.Request)

	AddLike(w http.ResponseWriter, r *http.Request)

	DeleteLike(w http.ResponseWriter, r *http.Request)
}

type postControllerImpl struct {
	serv domain.PostService
}

func (con postControllerImpl) AddLike(w http.ResponseWriter, r *http.Request) {
	userID, err := delivery.ParseUintParam(r, "userid")
	if err != nil {
		delivery.WriteResponse(w, http.StatusBadRequest, "Invalid userID provided")
		return
	}
	postID, err := delivery.ParseUintParam(r, "postid")
	if err != nil {
		delivery.WriteResponse(w, http.StatusBadRequest, "Invalid postID provided")
		return
	}

	err = con.serv.AddLike(userID, postID)
	if err == service.ErrAlreadyExisting {
		delivery.WriteResponse(w, http.StatusConflict, "Like already exists")
		return
	}
	if err == service.ErrIncorrectParameters {
		delivery.WriteResponse(w, http.StatusBadRequest, "Incorrect parameters provided ")
		return
	}
	if err == service.ErrDependencyNotSatisfied {
		delivery.WriteResponse(w, http.StatusNotFound, "User or Post doesn't exist")
		return
	}
	if err != nil {
		delivery.WriteResponse(w, http.StatusInternalServerError, "")
		return
	}

	delivery.WriteResponse(w, http.StatusCreated, "Like created successfully")
}

func (con postControllerImpl) Create(w http.ResponseWriter, r *http.Request) {
	userID, err := delivery.ParseUintParam(r, "userid")
	if err != nil {
		delivery.WriteResponse(w, http.StatusBadRequest, "Invalid ID provided")
		return
	}
	var createReq struct {
		Title       string `json:"Title"`
		Description string `json:"Description"`
		Content     string `json:"Content"`
	}
	err = delivery.ReadJSONRequest(r, &createReq)
	if err != nil {
		delivery.WriteResponse(w, http.StatusBadRequest, "Incorrect request format")
		return
	}

	id, err := con.serv.Create(userID, createReq.Title, createReq.Description, createReq.Content)
	if err == service.ErrIncorrectParameters {
		delivery.WriteResponse(w, http.StatusBadRequest, "Invalid provided parameters")
		return
	}
	if err == service.ErrDependencyNotSatisfied {
		delivery.WriteResponse(w, http.StatusNotFound, "Unexistent user")
		return
	}
	if err == service.ErrAlreadyExisting {
		delivery.WriteResponse(w, http.StatusConflict, "Repeated title")
		return
	}
	if err != nil {
		delivery.WriteResponse(w, http.StatusInternalServerError, "")
		return
	}

	response := struct {
		ID uint `json:"ID"`
	}{
		ID: id,
	}
	delivery.WriteJSONResponse(w, http.StatusCreated, response)
}

func (con postControllerImpl) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := delivery.ParseUintParam(r, "postid")
	if err != nil {
		delivery.WriteResponse(w, http.StatusBadRequest, "Invalid ID provided")
		return
	}

	err = con.serv.Delete(id)
	if err == service.ErrIncorrectParameters {
		delivery.WriteResponse(w, http.StatusBadRequest, "Invalid parameters provided")
		return
	}
	if err == service.ErrNotExistingEntity {
		delivery.WriteResponse(w, http.StatusNotFound, "Post doens't exist")
		return
	}
	if err != nil {
		delivery.WriteResponse(w, http.StatusInternalServerError, "")
		return
	}

	delivery.WriteResponse(w, http.StatusOK, "Post deleted successfully")
}

func (con postControllerImpl) DeleteLike(w http.ResponseWriter, r *http.Request) {
	userID, err := delivery.ParseUintParam(r, "userid")
	if err != nil {
		delivery.WriteResponse(w, http.StatusBadRequest, "Invalid userID provided")
		return
	}
	postID, err := delivery.ParseUintParam(r, "postid")
	if err != nil {
		delivery.WriteResponse(w, http.StatusBadRequest, "Invalid postID provided")
		return
	}

	err = con.serv.DeleteLike(userID, postID)
	if err == service.ErrIncorrectParameters {
		delivery.WriteResponse(w, http.StatusBadRequest, "Invalid parameters provided")
		return
	}
	if err == service.ErrNotExistingEntity {
		delivery.WriteResponse(w, http.StatusNotFound, "Like doesn't exist")
		return
	}
	if err != nil {
		delivery.WriteResponse(w, http.StatusInternalServerError, "")
		return
	}

	delivery.WriteResponse(w, http.StatusOK, "Like deleted successfully")
}

func (con postControllerImpl) GetByID(w http.ResponseWriter, r *http.Request) {
	postID, err := delivery.ParseUintParam(r, "postid")
	if err != nil {
		delivery.WriteResponse(w, http.StatusBadRequest, "Invalid ID provided ")
		return
	}

	post, err := con.serv.GetByID(postID)
	if err == service.ErrIncorrectParameters {
		delivery.WriteResponse(w, http.StatusBadRequest, "Invalid parameters provided")
		return
	}
	if err == service.ErrNotExistingEntity {
		delivery.WriteResponse(w, http.StatusNotFound, "Post doesn't exist")
		return
	}
	if err != nil {
		delivery.WriteResponse(w, http.StatusInternalServerError, "")
		return
	}

	delivery.WriteJSONResponse(w, http.StatusOK, post)
}

func (con postControllerImpl) GetByUser(w http.ResponseWriter, r *http.Request) {
	id, err := delivery.ParseUintParam(r, "userid")
	if err != nil {
		delivery.WriteResponse(w, http.StatusBadRequest, "Invalid provided ID")
		return
	}

	posts, err := con.serv.GetByUser(id)
	if err == service.ErrIncorrectParameters {
		delivery.WriteResponse(w, http.StatusBadRequest, "Invalid parameters provided")
		return
	}
	if err == service.ErrNotExistingEntity {
		delivery.WriteResponse(w, http.StatusNotFound, "No post found")
		return
	}
	if err != nil {
		delivery.WriteResponse(w, http.StatusInternalServerError, "")
		return
	}

	delivery.WriteJSONResponse(w, http.StatusOK, posts)
}

func (con postControllerImpl) GetPopularAllTime(w http.ResponseWriter, r *http.Request) {
	posts, err := con.serv.GetPopularAllTime()
	if err == service.ErrNotExistingEntity {
		delivery.WriteResponse(w, http.StatusNotFound, "No post found")
		return
	}
	if err != nil {
		delivery.WriteResponse(w, http.StatusInternalServerError, "")
		return
	}

	delivery.WriteJSONResponse(w, http.StatusOK, posts)
}

func (con postControllerImpl) GetPopularLastMonth(w http.ResponseWriter, r *http.Request) {
	posts, err := con.serv.GetPopularLastMonth()
	if err == service.ErrNotExistingEntity {
		delivery.WriteResponse(w, http.StatusNotFound, "No post found")
		return
	}
	if err != nil {
		delivery.WriteResponse(w, http.StatusInternalServerError, "")
		return
	}

	delivery.WriteJSONResponse(w, http.StatusOK, posts)
}

func (con postControllerImpl) GetPopularLastWeek(w http.ResponseWriter, r *http.Request) {
	posts, err := con.serv.GetPopularLastWeek()
	if err == service.ErrNotExistingEntity {
		delivery.WriteResponse(w, http.StatusNotFound, "No post found")
		return
	}
	if err != nil {
		delivery.WriteResponse(w, http.StatusInternalServerError, "")
		return
	}

	delivery.WriteJSONResponse(w, http.StatusOK, posts)
}

func (con postControllerImpl) GetPopularToday(w http.ResponseWriter, r *http.Request) {
	posts, err := con.serv.GetPopularToday()
	if err == service.ErrNotExistingEntity {
		delivery.WriteResponse(w, http.StatusNotFound, "No post found")
		return
	}
	if err != nil {
		delivery.WriteResponse(w, http.StatusInternalServerError, "")
		return
	}

	delivery.WriteJSONResponse(w, http.StatusOK, posts)
}

func (con postControllerImpl) UpdateContent(w http.ResponseWriter, r *http.Request) {
	postID, err := delivery.ParseUintParam(r, "postid")
	if err != nil {
		delivery.WriteResponse(w, http.StatusBadRequest, "Invalid provided ID")
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

	err = con.serv.UpdateContent(postID, updateReq.Content)
	if err == service.ErrIncorrectParameters {
		delivery.WriteResponse(w, http.StatusBadRequest, "Invalid parameters provided")
		return
	}
	if err == service.ErrNotExistingEntity {
		delivery.WriteResponse(w, http.StatusNotFound, "Post doesn't exist")
		return
	}
	if err != nil {
		delivery.WriteResponse(w, http.StatusInternalServerError, "")
		return
	}

	delivery.WriteResponse(w, http.StatusOK, "Post updated succesfully")
}

func (con postControllerImpl) UpdateDescription(w http.ResponseWriter, r *http.Request) {
	postID, err := delivery.ParseUintParam(r, "postid")
	if err != nil {
		delivery.WriteResponse(w, http.StatusBadRequest, "Invalid provided ID")
		return
	}

	var updateReq struct {
		Description string `json:"Description"`
	}
	err = delivery.ReadJSONRequest(r, &updateReq)
	if err != nil {
		delivery.WriteResponse(w, http.StatusBadRequest, "Incorrect request format")
		return
	}

	err = con.serv.UpdateDescription(postID, updateReq.Description)
	if err == service.ErrIncorrectParameters {
		delivery.WriteResponse(w, http.StatusBadRequest, "Invalid parameters provided")
		return
	}
	if err == service.ErrNotExistingEntity {
		delivery.WriteResponse(w, http.StatusNotFound, "Post doesn't exist")
		return
	}
	if err != nil {
		delivery.WriteResponse(w, http.StatusInternalServerError, "")
		return
	}

	delivery.WriteResponse(w, http.StatusOK, "Post updated succesfully")
}

func (con postControllerImpl) UpdateTitle(w http.ResponseWriter, r *http.Request) {
	postID, err := delivery.ParseUintParam(r, "postid")
	if err != nil {
		delivery.WriteResponse(w, http.StatusBadRequest, "Invalid provided ID")
		return
	}

	var updateReq struct {
		Title string `json:"Title"`
	}
	err = delivery.ReadJSONRequest(r, &updateReq)
	if err != nil {
		delivery.WriteResponse(w, http.StatusBadRequest, "Incorrect request format")
		return
	}

	err = con.serv.UpdateTitle(postID, updateReq.Title)
	if err == service.ErrIncorrectParameters {
		delivery.WriteResponse(w, http.StatusBadRequest, "Invalid parameters provided")
		return
	}
	if err == service.ErrNotExistingEntity {
		delivery.WriteResponse(w, http.StatusNotFound, "Post doesn't exist")
		return
	}
	if err == service.ErrAlreadyExisting {
		delivery.WriteResponse(w, http.StatusConflict, "Repeated title")
		return
	}
	if err != nil {
		delivery.WriteResponse(w, http.StatusInternalServerError, "")
		return
	}

	delivery.WriteResponse(w, http.StatusOK, "Post updated succesfully")
}

func NewPostController(serv domain.PostService) PostController {
	return postControllerImpl{serv: serv}
}

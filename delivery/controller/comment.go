package controller

import (
	"net/http"

	"github.com/AlejandroJorge/forum-rest-api/delivery"
	"github.com/AlejandroJorge/forum-rest-api/domain"
	"github.com/AlejandroJorge/forum-rest-api/service"
)

type CommentController interface {
	Create(w http.ResponseWriter, r *http.Request)

	Delete(w http.ResponseWriter, r *http.Request)

	UpdateContent(w http.ResponseWriter, r *http.Request)

	GetByID(w http.ResponseWriter, r *http.Request)

	GetByPost(w http.ResponseWriter, r *http.Request)

	GetByUser(w http.ResponseWriter, r *http.Request)

	AddLike(w http.ResponseWriter, r *http.Request)

	DeleteLike(w http.ResponseWriter, r *http.Request)
}

type commentControllerImpl struct {
	serv domain.CommentService
}

func (con commentControllerImpl) AddLike(w http.ResponseWriter, r *http.Request) {
	userID, err := delivery.ParseUintParam(r, "userid")
	if err != nil {
		delivery.WriteResponse(w, http.StatusBadRequest, "Invalid userID provided")
		return
	}

	commentID, err := delivery.ParseUintParam(r, "commentid")
	if err != nil {
		delivery.WriteResponse(w, http.StatusBadRequest, "Invalid commentID provided")
		return
	}

	err = con.serv.AddLike(userID, commentID)
	if err == service.ErrIncorrectParameters {
		delivery.WriteResponse(w, http.StatusBadRequest, "Invalid parameters provided")
		return
	}
	if err == service.ErrDependencyNotSatisfied {
		delivery.WriteResponse(w, http.StatusNotFound, "User or Post doesn't exist")
	}
	if err != nil {
		delivery.WriteResponse(w, http.StatusInternalServerError, "")
		return
	}

	delivery.WriteResponse(w, http.StatusOK, "Like created successfully")
}

func (con commentControllerImpl) Create(w http.ResponseWriter, r *http.Request) {
	userID, err := delivery.ParseUintParam(r, "userid")
	var createReq struct {
		PostID  uint   `json:"PostId"`
		Content string `json:"Content"`
	}
	err = delivery.ReadJSONRequest(r, &createReq)
	if err != nil {
		delivery.WriteResponse(w, http.StatusBadRequest, "Incorrect request format")
		return
	}

	id, err := con.serv.Create(userID, createReq.PostID, createReq.Content)
	if err == service.ErrIncorrectParameters {
		delivery.WriteResponse(w, http.StatusBadRequest, "Invalid parameters provided")
		return
	}
	if err == service.ErrDependencyNotSatisfied {
		delivery.WriteResponse(w, http.StatusBadRequest, "User or Post doesn't exist")
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

func (con commentControllerImpl) Delete(w http.ResponseWriter, r *http.Request) {
	userID, err := delivery.ParseUintParam(r, "userid")
	if err != nil {
		delivery.WriteResponse(w, http.StatusOK, "Invalid ID provided")
		return
	}

	err = con.serv.Delete(userID)
	if err == service.ErrIncorrectParameters {
		delivery.WriteResponse(w, http.StatusBadRequest, "Invalid parameters provided")
		return
	}
	if err == service.ErrNotExistingEntity {
		delivery.WriteResponse(w, http.StatusNotFound, "Comment doesn't exist")
		return
	}
	if err != nil {
		delivery.WriteResponse(w, http.StatusInternalServerError, "")
		return
	}

	delivery.WriteResponse(w, http.StatusOK, "Post deleted successfully")
}

func (con commentControllerImpl) DeleteLike(w http.ResponseWriter, r *http.Request) {
	userID, err := delivery.ParseUintParam(r, "userid")
	if err != nil {
		delivery.WriteResponse(w, http.StatusBadRequest, "Invalid userID provided")
		return
	}

	commentID, err := delivery.ParseUintParam(r, "commentid")
	if err != nil {
		delivery.WriteResponse(w, http.StatusBadRequest, "Invalid postID provided")
		return
	}

	err = con.serv.DeleteLike(userID, commentID)
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

func (con commentControllerImpl) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := delivery.ParseUintParam(r, "commentid")
	if err != nil {
		delivery.WriteResponse(w, http.StatusOK, "Invalid ID provided")
		return
	}

	comment, err := con.serv.GetByID(id)
	if err == service.ErrIncorrectParameters {
		delivery.WriteResponse(w, http.StatusBadRequest, "Invalid parameters provided")
		return
	}
	if err == service.ErrNotExistingEntity {
		delivery.WriteResponse(w, http.StatusNotFound, "Comment doesn't exist")
		return
	}
	if err != nil {
		delivery.WriteResponse(w, http.StatusInternalServerError, "")
		return
	}

	delivery.WriteJSONResponse(w, http.StatusOK, comment)
}

func (con commentControllerImpl) GetByPost(w http.ResponseWriter, r *http.Request) {
	id, err := delivery.ParseUintParam(r, "postid")
	if err != nil {
		delivery.WriteResponse(w, http.StatusOK, "Invalid ID provided")
		return
	}

	posts, err := con.serv.GetByPost(id)
	if err == service.ErrIncorrectParameters {
		delivery.WriteResponse(w, http.StatusBadRequest, "Invalid parameters provided")
		return
	}
	if err == service.ErrNotExistingEntity {
		delivery.WriteResponse(w, http.StatusNotFound, "No posts found")
		return
	}
	if err != nil {
		delivery.WriteResponse(w, http.StatusInternalServerError, "")
		return
	}

	delivery.WriteJSONResponse(w, http.StatusOK, posts)
}

func (con commentControllerImpl) GetByUser(w http.ResponseWriter, r *http.Request) {
	id, err := delivery.ParseUintParam(r, "userid")
	if err != nil {
		delivery.WriteResponse(w, http.StatusOK, "Invalid ID provided")
		return
	}

	posts, err := con.serv.GetByUser(id)
	if err == service.ErrIncorrectParameters {
		delivery.WriteResponse(w, http.StatusBadRequest, "Invalid parameters provided")
		return
	}
	if err == service.ErrNotExistingEntity {
		delivery.WriteResponse(w, http.StatusNotFound, "No posts found")
		return
	}
	if err != nil {
		delivery.WriteResponse(w, http.StatusInternalServerError, "")
		return
	}

	delivery.WriteJSONResponse(w, http.StatusOK, posts)
}

func (con commentControllerImpl) UpdateContent(w http.ResponseWriter, r *http.Request) {
	id, err := delivery.ParseUintParam(r, "commentid")
	if err != nil {
		delivery.WriteResponse(w, http.StatusBadRequest, "Invalid ID provided")
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

	err = con.serv.Update(id, updateReq.Content)
	if err == service.ErrIncorrectParameters {
		delivery.WriteResponse(w, http.StatusBadRequest, "Invalid parameters provided")
		return
	}
	if err == service.ErrNotExistingEntity {
		delivery.WriteResponse(w, http.StatusNotFound, "Comment doesn't exist")
		return
	}
	if err != nil {
		delivery.WriteResponse(w, http.StatusInternalServerError, "")
		return
	}

	delivery.WriteResponse(w, http.StatusOK, "Comment updated successfully")
}

func NewCommentController(serv domain.CommentService) CommentController {
	return commentControllerImpl{serv: serv}
}

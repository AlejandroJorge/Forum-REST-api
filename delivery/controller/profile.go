package controller

import (
	"net/http"

	"github.com/AlejandroJorge/forum-rest-api/delivery"
	"github.com/AlejandroJorge/forum-rest-api/domain"
	"github.com/AlejandroJorge/forum-rest-api/service"
)

type ProfileController interface {
	Create(w http.ResponseWriter, r *http.Request)

	Delete(w http.ResponseWriter, r *http.Request)

	UpdateTagName(w http.ResponseWriter, r *http.Request)

	UpdateDisplayName(w http.ResponseWriter, r *http.Request)

	UpdatePicturePath(w http.ResponseWriter, r *http.Request)

	UpdateBackgroundPath(w http.ResponseWriter, r *http.Request)

	GetByUserID(w http.ResponseWriter, r *http.Request)

	GetByTagName(w http.ResponseWriter, r *http.Request)

	GetFollowersByID(w http.ResponseWriter, r *http.Request)

	GetFollowersByTagName(w http.ResponseWriter, r *http.Request)

	GetFollowsByID(w http.ResponseWriter, r *http.Request)

	GetFollowsByTagName(w http.ResponseWriter, r *http.Request)

	AddFollow(w http.ResponseWriter, r *http.Request)

	DeleteFollow(w http.ResponseWriter, r *http.Request)
}

type profileControllerImpl struct {
	serv domain.ProfileService
}

func (con profileControllerImpl) AddFollow(w http.ResponseWriter, r *http.Request) {
	userID, err := delivery.ParseUintParam(r, "userid")
	if err != nil {
		delivery.WriteResponse(w, http.StatusBadRequest, "Invalid userID provided")
		return
	}

	followedID, err := delivery.ParseUintParam(r, "followedid")
	if err != nil {
		delivery.WriteResponse(w, http.StatusBadRequest, "Invalid userID provided")
		return
	}

	err = con.serv.AddFollow(userID, followedID)
	if err == service.ErrAlreadyExisting {
		delivery.WriteResponse(w, http.StatusConflict, "This follow already exists")
		return
	}
	if err == service.ErrDependencyNotSatisfied {
		delivery.WriteResponse(w, http.StatusBadRequest, "Follower or followed doesn't exist")
		return
	}
	if err == service.ErrIncorrectParameters {
		delivery.WriteResponse(w, http.StatusBadRequest, "Incorrect parameters provided")
		return
	}
	if err != nil {
		delivery.WriteResponse(w, http.StatusInternalServerError, "")
		return
	}

	delivery.WriteResponse(w, http.StatusCreated, "Follow successfully created")
}

func (con profileControllerImpl) Create(w http.ResponseWriter, r *http.Request) {
	userID, err := delivery.ParseUintParam(r, "userid")
	if err != nil {
		delivery.WriteResponse(w, http.StatusBadRequest, "Invalid ID provided")
		return
	}

	var createReq struct {
		DisplayName string `json:"DisplayName"`
		TagName     string `json:"TagName"`
	}
	err = delivery.ReadJSONRequest(r, &createReq)
	if err != nil {
		delivery.WriteResponse(w, http.StatusBadRequest, "Incorrect request format")
		return
	}

	id, err := con.serv.Create(userID, createReq.TagName, createReq.DisplayName)
	if err == service.ErrDependencyNotSatisfied {
		delivery.WriteResponse(w, http.StatusBadRequest, "User doesn't exist")
		return
	}
	if err == service.ErrProfileExistsOrTagNameIsRepeated {
		delivery.WriteResponse(w, http.StatusBadRequest, "Profile already exists for this user or tag name is repeated")
		return
	}
	if err == service.ErrIncorrectParameters {
		delivery.WriteResponse(w, http.StatusBadRequest, "Incorrect parameters provided")
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

func (con profileControllerImpl) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := delivery.ParseUintParam(r, "userid")
	if err != nil {
		delivery.WriteResponse(w, http.StatusBadRequest, "Invalid id provided")
		return
	}

	err = con.serv.Delete(id)
	if err == service.ErrIncorrectParameters {
		delivery.WriteResponse(w, http.StatusBadRequest, "Invalid parameters provided")
		return
	}
	if err == service.ErrNotExistingEntity {
		delivery.WriteResponse(w, http.StatusNotFound, "Profile doesn't exist")
		return
	}

	delivery.WriteResponse(w, http.StatusOK, "Profile deleted successfully")
}

func (con profileControllerImpl) DeleteFollow(w http.ResponseWriter, r *http.Request) {
	userID, err := delivery.ParseUintParam(r, "userid")
	if err != nil {
		delivery.WriteResponse(w, http.StatusBadRequest, "Invalid ID provided")
	}

	followedID, err := delivery.ParseUintParam(r, "followedid")
	if err != nil {
		delivery.WriteResponse(w, http.StatusBadRequest, "Invalid ID provided")
	}

	err = con.serv.DeleteFollow(userID, followedID)
	if err == service.ErrIncorrectParameters {
		delivery.WriteResponse(w, http.StatusBadRequest, "Incorrect parameters provided")
		return
	}
	if err == service.ErrNotExistingEntity {
		delivery.WriteResponse(w, http.StatusNotFound, "Follower or followed doesn't exist")
		return
	}

	delivery.WriteResponse(w, http.StatusOK, "Follow deleted successfully")
}

func (con profileControllerImpl) GetByTagName(w http.ResponseWriter, r *http.Request) {
	tagName, err := delivery.ParseStringParam(r, "tagname")
	if err != nil {
		delivery.WriteResponse(w, http.StatusBadRequest, "Invalid tagname provided")
		return
	}

	profile, err := con.serv.GetByTagName(tagName)
	if err == service.ErrIncorrectParameters {
		delivery.WriteResponse(w, http.StatusBadRequest, "Incorrect parameters provided")
		return
	}
	if err == service.ErrNotExistingEntity {
		delivery.WriteResponse(w, http.StatusNotFound, "Profile doesn't exist")
		return
	}
	if err != nil {
		delivery.WriteResponse(w, http.StatusInternalServerError, "")
		return
	}

	delivery.WriteJSONResponse(w, http.StatusOK, profile)
}

func (con profileControllerImpl) GetByUserID(w http.ResponseWriter, r *http.Request) {
	id, err := delivery.ParseUintParam(r, "userid")
	if err != nil {
		delivery.WriteResponse(w, http.StatusBadRequest, "Invalid id provided")
		return
	}

	profile, err := con.serv.GetByUserID(id)
	if err == service.ErrIncorrectParameters {
		delivery.WriteResponse(w, http.StatusBadRequest, "Incorrect parameters provided")
		return
	}
	if err == service.ErrNotExistingEntity {
		delivery.WriteResponse(w, http.StatusNotFound, "Profile doesn't exist")
		return
	}
	if err != nil {
		delivery.WriteResponse(w, http.StatusInternalServerError, "")
		return
	}

	delivery.WriteJSONResponse(w, http.StatusOK, profile)
}

func (con profileControllerImpl) GetFollowersByID(w http.ResponseWriter, r *http.Request) {
	id, err := delivery.ParseUintParam(r, "userid")
	if err != nil {
		delivery.WriteResponse(w, http.StatusBadRequest, "Invalid id provided")
		return
	}

	followers, err := con.serv.GetFollowersByID(id)
	if err == service.ErrIncorrectParameters {
		delivery.WriteResponse(w, http.StatusBadRequest, "Incorrect parameters provided")
		return
	}
	if err == service.ErrNotExistingEntity {
		delivery.WriteResponse(w, http.StatusNotFound, "No followers found")
		return
	}
	if err != nil {
		delivery.WriteResponse(w, http.StatusInternalServerError, "")
		return
	}

	delivery.WriteJSONResponse(w, http.StatusOK, followers)
}

func (con profileControllerImpl) GetFollowersByTagName(w http.ResponseWriter, r *http.Request) {
	tagName, err := delivery.ParseStringParam(r, "tagname")
	if err != nil {
		delivery.WriteResponse(w, http.StatusBadRequest, "Invalid tagname provided")
		return
	}

	followers, err := con.serv.GetFollowersByTagName(tagName)
	if err == service.ErrIncorrectParameters {
		delivery.WriteResponse(w, http.StatusBadRequest, "Incorrect parameters provided")
		return
	}
	if err == service.ErrNotExistingEntity {
		delivery.WriteResponse(w, http.StatusNotFound, "No followers found")
		return
	}
	if err != nil {
		delivery.WriteResponse(w, http.StatusInternalServerError, "")
		return
	}

	delivery.WriteJSONResponse(w, http.StatusOK, followers)
}

func (con profileControllerImpl) GetFollowsByID(w http.ResponseWriter, r *http.Request) {
	id, err := delivery.ParseUintParam(r, "userid")
	if err != nil {
		delivery.WriteResponse(w, http.StatusBadRequest, "Invalid tagname provided")
		return
	}

	follows, err := con.serv.GetFollowsByID(id)
	if err == service.ErrIncorrectParameters {
		delivery.WriteResponse(w, http.StatusBadRequest, "Incorrect parameters provided")
		return
	}
	if err == service.ErrNotExistingEntity {
		delivery.WriteResponse(w, http.StatusNotFound, "No follows found")
		return
	}
	if err != nil {
		delivery.WriteResponse(w, http.StatusInternalServerError, "")
		return
	}

	delivery.WriteJSONResponse(w, http.StatusOK, follows)
}

func (con profileControllerImpl) GetFollowsByTagName(w http.ResponseWriter, r *http.Request) {
	tagName, err := delivery.ParseStringParam(r, "tagname")
	if err != nil {
		delivery.WriteResponse(w, http.StatusBadRequest, "Invalid tagname provided")
		return
	}

	follows, err := con.serv.GetFollowsByTagName(tagName)
	if err == service.ErrIncorrectParameters {
		delivery.WriteResponse(w, http.StatusBadRequest, "Incorrect parameters provided")
		return
	}
	if err == service.ErrNotExistingEntity {
		delivery.WriteResponse(w, http.StatusNotFound, "No follows found")
		return
	}
	if err != nil {
		delivery.WriteResponse(w, http.StatusInternalServerError, "")
		return
	}

	delivery.WriteJSONResponse(w, http.StatusOK, follows)
}

func (con profileControllerImpl) UpdateBackgroundPath(w http.ResponseWriter, r *http.Request) {
	id, err := delivery.ParseUintParam(r, "userid")
	if err != nil {
		delivery.WriteResponse(w, http.StatusBadRequest, "Invalid id provided")
		return
	}

	var updateReq struct {
		BackgroundPath string `json:"BackgroundPath"`
	}
	err = delivery.ReadJSONRequest(r, &updateReq)
	if err != nil {
		delivery.WriteResponse(w, http.StatusBadRequest, "Incorrect request format")
		return
	}

	err = con.serv.UpdateBackgroundPath(id, updateReq.BackgroundPath)
	if err == service.ErrNotExistingEntity {
		delivery.WriteResponse(w, http.StatusNotFound, "Profile with that ID doesn't exist")
		return
	}
	if err != nil {
		delivery.WriteResponse(w, http.StatusInternalServerError, "")
		return
	}

	delivery.WriteResponse(w, http.StatusOK, "Profile updated successfully")
}

func (con profileControllerImpl) UpdateDisplayName(w http.ResponseWriter, r *http.Request) {
	id, err := delivery.ParseUintParam(r, "userid")
	if err != nil {
		delivery.WriteResponse(w, http.StatusBadRequest, "Invalid id provided")
		return
	}

	var updateReq struct {
		DisplayName string `json:"DisplayName"`
	}
	err = delivery.ReadJSONRequest(r, &updateReq)
	if err != nil {
		delivery.WriteResponse(w, http.StatusBadRequest, "Incorrect request format")
		return
	}

	err = con.serv.UpdateDisplayName(id, updateReq.DisplayName)
	if err == service.ErrNotExistingEntity {
		delivery.WriteResponse(w, http.StatusNotFound, "Profile with that ID doesn't exist")
		return
	}
	if err != nil {
		delivery.WriteResponse(w, http.StatusInternalServerError, "")
		return
	}

	delivery.WriteResponse(w, http.StatusOK, "Profile updated successfully")
}

func (con profileControllerImpl) UpdatePicturePath(w http.ResponseWriter, r *http.Request) {
	id, err := delivery.ParseUintParam(r, "userid")
	if err != nil {
		delivery.WriteResponse(w, http.StatusBadRequest, "Invalid id provided")
		return
	}

	var updateReq struct {
		PicturePath string `json:"PicturePath"`
	}
	err = delivery.ReadJSONRequest(r, &updateReq)
	if err != nil {
		delivery.WriteResponse(w, http.StatusBadRequest, "Incorrect request format")
		return
	}

	err = con.serv.UpdatePicturePath(id, updateReq.PicturePath)
	if err == service.ErrNotExistingEntity {
		delivery.WriteResponse(w, http.StatusNotFound, "Profile with that ID doesn't exist")
		return
	}
	if err != nil {
		delivery.WriteResponse(w, http.StatusInternalServerError, "")
		return
	}

	delivery.WriteResponse(w, http.StatusOK, "Profile updated successfully")
}

func (con profileControllerImpl) UpdateTagName(w http.ResponseWriter, r *http.Request) {
	id, err := delivery.ParseUintParam(r, "userid")
	if err != nil {
		delivery.WriteResponse(w, http.StatusBadRequest, "Invalid id provided")
		return
	}

	var updateReq struct {
		TagName string `json:"TagName"`
	}
	err = delivery.ReadJSONRequest(r, &updateReq)
	if err != nil {
		delivery.WriteResponse(w, http.StatusBadRequest, "Incorrect request format")
		return
	}

	err = con.serv.UpdateTagName(id, updateReq.TagName)
	if err == service.ErrNotExistingEntity {
		delivery.WriteResponse(w, http.StatusNotFound, "Profile with that ID doesn't exist")
		return
	}
	if err == service.ErrAlreadyExisting {
		delivery.WriteResponse(w, http.StatusConflict, "That tagname already exists")
		return
	}
	if err != nil {
		delivery.WriteResponse(w, http.StatusInternalServerError, "")
		return
	}

	delivery.WriteResponse(w, http.StatusOK, "Profile updated successfully")
}

func NewProfileController(serv domain.ProfileService) ProfileController {
	return profileControllerImpl{serv: serv}
}

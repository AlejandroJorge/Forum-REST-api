package controller

import (
	"net/http"
	"strconv"

	"github.com/AlejandroJorge/forum-rest-api/delivery"
	"github.com/AlejandroJorge/forum-rest-api/domain"
	"github.com/AlejandroJorge/forum-rest-api/util"
	"github.com/gorilla/mux"
)

type profileController struct {
	serv domain.ProfileService
}

func (con profileController) GetByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idStr, ok := params["id"]
	if !ok {
		delivery.WriteResponse(w, http.StatusBadRequest, "No provided ID")
		return
	}

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		delivery.WriteResponse(w, http.StatusBadRequest, "Id provided isn't a number")
		return
	}

	profile, err := con.serv.GetByUserID(uint(id))
	if err == util.ErrNoCorrespondingUser {
		delivery.WriteResponse(w, http.StatusBadRequest, "There's no user corresponding to this id")
	}
	if err == util.ErrEmptySelection {
		delivery.WriteResponse(w, http.StatusNotFound, "This user doesn't have a profile")
		return
	}
	if err != nil {
		delivery.WriteResponse(w, http.StatusInternalServerError, "Couldn't retrieve profile")
		return
	}

	delivery.WriteJSONResponse(w, http.StatusOK, profile)
}

func (con profileController) GetByTagName(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	tagName, ok := params["tagname"]
	if !ok {
		delivery.WriteResponse(w, http.StatusBadRequest, "No provided tagname")
		return
	}

	profile, err := con.serv.GetByTagName(tagName)
	if err == util.ErrEmptySelection {
		delivery.WriteResponse(w, http.StatusNotFound, "There was no profile with this tagName")
		return
	}
	if err != nil {
		delivery.WriteResponse(w, http.StatusInternalServerError, "Couldn't retrieve profile")
		return
	}

	delivery.WriteJSONResponse(w, http.StatusOK, profile)
}

func (con profileController) GetFollowsByID(w http.ResponseWriter, r *http.Request) {
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

	follows, err := con.serv.GetFollowsByID(uint(id))
	if err == util.ErrEmptySelection {
		delivery.WriteResponse(w, http.StatusNotFound, "No follows found")
		return
	}
	if err != nil {
		delivery.WriteResponse(w, http.StatusInternalServerError, "Couldn't retrieve follows")
		return
	}

	delivery.WriteJSONResponse(w, http.StatusOK, follows)
}

func (con profileController) GetFollowsByTagName(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	tagName, ok := params["tagname"]
	if !ok {
		delivery.WriteResponse(w, http.StatusBadRequest, "No provided tagname")
		return
	}

	follows, err := con.serv.GetFollowsByTagName(tagName)
	if err == util.ErrEmptySelection {
		delivery.WriteResponse(w, http.StatusNotFound, "No follows found")
		return
	}
	if err != nil {
		delivery.WriteResponse(w, http.StatusInternalServerError, "Couldn't retrieve follows")
		return
	}

	delivery.WriteJSONResponse(w, http.StatusOK, follows)
}

func (con profileController) GetFollowersByID(w http.ResponseWriter, r *http.Request) {
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

	followers, err := con.serv.GetFollowersByID(uint(id))
	if err == util.ErrEmptySelection {
		delivery.WriteResponse(w, http.StatusNotFound, "No followers found")
		return
	}
	if err != nil {
		delivery.WriteResponse(w, http.StatusInternalServerError, "Couldn't retrieve followers")
		return
	}

	delivery.WriteJSONResponse(w, http.StatusOK, followers)
}

func (con profileController) GetFollowersByTagName(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	tagName, ok := params["tagname"]
	if !ok {
		delivery.WriteResponse(w, http.StatusBadRequest, "No provided tagname")
		return
	}

	followers, err := con.serv.GetFollowersByTagName(tagName)
	if err == util.ErrEmptySelection {
		delivery.WriteResponse(w, http.StatusNotFound, "No followers found")
		return
	}
	if err != nil {
		delivery.WriteResponse(w, http.StatusInternalServerError, "Couldn't retrieve followers")
		return
	}

	delivery.WriteJSONResponse(w, http.StatusOK, followers)
}

func (con profileController) Create(w http.ResponseWriter, r *http.Request) {
	var createReq struct {
		UserID         uint   `json:"UserID"`
		DisplayName    string `json:"DisplayName"`
		TagName        string `json:"TagName"`
		PicturePath    string `json:"PicturePath,omitempty"`
		BackgroundPath string `json:"BackgroundPath,omitempty"`
	}
	err := delivery.ReadJSONRequest(r, &createReq)
	if err != nil {
		delivery.WriteResponse(w, http.StatusBadRequest, "Incorrect format of request")
		return
	}

	id, err := con.serv.CreateNew(struct {
		UserID         uint
		DisplayName    string
		TagName        string
		PicturePath    string
		BackgroundPath string
	}{
		UserID:         createReq.UserID,
		DisplayName:    createReq.DisplayName,
		TagName:        createReq.TagName,
		PicturePath:    createReq.PicturePath,
		BackgroundPath: createReq.BackgroundPath,
	})
	if err == util.ErrRepeatedEntity {
		delivery.WriteResponse(w, http.StatusConflict, "There's already a profile for the same User or TagName is taken")
		return
	}
	if err == util.ErrNoCorrespondingUser {
		delivery.WriteResponse(w, http.StatusBadRequest, "There's no user corresponding to this ID")
		return
	}
	if err != nil {
		delivery.WriteResponse(w, http.StatusInternalServerError, "Couldn't create resource")
		return
	}

	response := struct {
		ID uint `json:"ID"`
	}{ID: id}
	delivery.WriteJSONResponse(w, http.StatusCreated, response)
}

func (con profileController) Update(w http.ResponseWriter, r *http.Request) {
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
		TagName        string `json:"TagName,omitempty"`
		DisplayName    string `json:"DisplayName,omitempty"`
		PicturePath    string `json:"PicturePath,omitempty"`
		BackgroundPath string `json:"BackgroundPath,omitempty"`
	}
	err = delivery.ReadJSONRequest(r, &updateReq)
	if err != nil {
		delivery.WriteResponse(w, http.StatusBadRequest, "Incorrect request format")
		return
	}

	err = con.serv.Update(uint(id), struct {
		UpdatedTagName        string
		UpdatedDisplayName    string
		UpdatedPicturePath    string
		UpdatedBackgroundPath string
	}{
		UpdatedTagName:        updateReq.TagName,
		UpdatedDisplayName:    updateReq.DisplayName,
		UpdatedPicturePath:    updateReq.PicturePath,
		UpdatedBackgroundPath: updateReq.BackgroundPath,
	})
	if err != nil {
		delivery.WriteResponse(w, http.StatusInternalServerError, "Couldn't update profile")
		return
	}

	delivery.WriteResponse(w, http.StatusOK, "Profile updated successfully")
}

func (con profileController) CreateFollow(w http.ResponseWriter, r *http.Request) {
	var followReq struct {
		FollowerID uint
		FollowedID uint
	}
	err := delivery.ReadJSONRequest(r, &followReq)
	if err != nil {
		delivery.WriteResponse(w, http.StatusBadRequest, "Incorrect request format")
		return
	}

	err = con.serv.AddFollow(followReq.FollowerID, followReq.FollowedID)
	if err != nil {
		delivery.WriteResponse(w, http.StatusInternalServerError, "Couldn't create follow relationship")
	}

	delivery.WriteResponse(w, http.StatusOK, "Registered follow successfully")
}

func (con profileController) DeleteFollow(w http.ResponseWriter, r *http.Request) {
	var followReq struct {
		FollowerID uint
		FollowedID uint
	}
	err := delivery.ReadJSONRequest(r, &followReq)
	if err != nil {
		delivery.WriteResponse(w, http.StatusBadRequest, "Incorrect request format")
		return
	}

	err = con.serv.DeleteFollow(followReq.FollowerID, followReq.FollowedID)
	if err != nil {
		delivery.WriteResponse(w, http.StatusInternalServerError, "Couldn't create follow relationship")
	}

	delivery.WriteResponse(w, http.StatusOK, "Deleted follow successfully")
}

func (con profileController) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idStr, ok := params["id"]
	if !ok {
		delivery.WriteResponse(w, http.StatusBadRequest, "No provided ID")
		return
	}

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		delivery.WriteResponse(w, http.StatusBadRequest, "Id provided isn't a number")
		return
	}

	err = con.serv.Delete(uint(id))
	if err != nil {
		delivery.WriteResponse(w, http.StatusInternalServerError, "Couldn't delete") // This is uncompressed, there's more errors to handle
		return
	}

	delivery.WriteResponse(w, http.StatusOK, "Deleted")
}

func NewProfileController(serv domain.ProfileService) profileController {
	return profileController{serv: serv}
}

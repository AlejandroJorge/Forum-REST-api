package controller

import (
	"net/http"
	"strconv"

	"github.com/AlejandroJorge/forum-rest-api/domain"
	"github.com/AlejandroJorge/forum-rest-api/util"
	"github.com/gorilla/mux"
)

type userController struct {
	serv domain.UserService
}

func NewUserController(serv domain.UserService) userController {
	return userController{serv: serv}
}

func (con userController) Create(w http.ResponseWriter, r *http.Request) {
	var createReq struct {
		NewEmail    string `json:"Email"`
		NewPassword string `json:"Password"`
	}
	err := util.ReadJSONRequest(r, &createReq)
	if err != nil {
		util.WriteResponse(w, http.StatusBadRequest, "Incorrect format of request")
		return
	}

	id, err := con.serv.CreateNew(struct {
		NewEmail    string
		NewPassword string
	}{
		NewEmail:    createReq.NewEmail,
		NewPassword: createReq.NewPassword,
	})
	if err == util.ErrRepeatedEntity {
		util.WriteResponse(w, http.StatusConflict, "Email already registered")
		return
	}
	if err == util.ErrIncorrectParameters {
		util.WriteResponse(w, http.StatusBadRequest, "Incorrect information provided")
		return
	}
	if err == util.ErrPasswordNotGenerated {
		util.WriteResponse(w, http.StatusInternalServerError, "Couldn't hash the password")
		return
	}
	if err != nil {
		util.WriteResponse(w, http.StatusInternalServerError, "")
		return
	}

	response := struct {
		ID uint `json:"ID"`
	}{ID: id}
	util.WriteJSONResponse(w, http.StatusCreated, response)
}

func (con userController) Get(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idStr, ok := params["id"]
	if !ok {
		util.WriteResponse(w, http.StatusBadRequest, "No provided ID")
		return
	}

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		util.WriteResponse(w, http.StatusBadRequest, "Id provided isn't a number")
		return
	}

	user, err := con.serv.GetByID(uint(id))
	if err == util.ErrEmptySelection {
		util.WriteResponse(w, http.StatusNotFound, "Nothing found")
		return
	}
	if err != nil {
		util.WriteResponse(w, http.StatusInternalServerError, "Something went wrong while retrieving the user")
		return
	}

	util.WriteJSONResponse(w, http.StatusOK, user)
}

func (con userController) Update(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idStr, ok := params["id"]
	if !ok {
		util.WriteResponse(w, http.StatusBadRequest, "No provided ID")
		return
	}

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		util.WriteResponse(w, http.StatusBadRequest, "Id provided isn't a number")
		return
	}

	var updateReq struct {
		Email    string `json:"Email"`
		Password string `json:"Password"`
	}
	err = util.ReadJSONRequest(r, &updateReq)
	if err != nil {
		util.WriteResponse(w, http.StatusBadRequest, "Incorrect format of request")
		return
	}

	err = con.serv.Update(uint(id), struct {
		UpdatedEmail    string
		UpdatedPassword string
	}{
		UpdatedEmail:    updateReq.Email,
		UpdatedPassword: updateReq.Password,
	})
	if err != nil {
		util.WriteResponse(w, http.StatusBadRequest, "Couldn't update") // This is uncompressed, there's more errors to handle
		return
	}

	util.WriteResponse(w, http.StatusOK, "Update successful")
}

func (con userController) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idStr, ok := params["id"]
	if !ok {
		util.WriteResponse(w, http.StatusBadRequest, "No provided ID")
		return
	}

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		util.WriteResponse(w, http.StatusBadRequest, "Id provided isn't a number")
		return
	}

	err = con.serv.Delete(uint(id))
	if err != nil {
		util.WriteResponse(w, http.StatusInternalServerError, "Couldn't delete") // This is uncompressed, there's more errors to handle
		return
	}

	util.WriteResponse(w, http.StatusOK, "Deleted")
}

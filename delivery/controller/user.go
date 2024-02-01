package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/AlejandroJorge/forum-rest-api/config"
	"github.com/AlejandroJorge/forum-rest-api/delivery"
	"github.com/AlejandroJorge/forum-rest-api/domain"
	"github.com/AlejandroJorge/forum-rest-api/service"
	"github.com/golang-jwt/jwt"
)

type UserController interface {
	Create(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	UpdateEmail(w http.ResponseWriter, r *http.Request)
	UpdatePassword(w http.ResponseWriter, r *http.Request)
	GetByID(w http.ResponseWriter, r *http.Request)
	CheckCredentials(w http.ResponseWriter, r *http.Request)
}

type userControllerImpl struct {
	serv domain.UserService
}

func NewUserController(serv domain.UserService) UserController {
	return userControllerImpl{serv: serv}
}

func (con userControllerImpl) Create(w http.ResponseWriter, r *http.Request) {
	var createReq struct {
		Email    string `json:"Email"`
		Password string `json:"Password"`
	}
	err := delivery.ReadJSONRequest(r, &createReq)
	if err != nil {
		fmt.Println("Error:", err)
		delivery.WriteResponse(w, http.StatusBadRequest, "Incorrect format of request")
		return
	}

	id, err := con.serv.Create(createReq.Email, createReq.Password)
	if err == service.ErrIncorrectParameters {
		delivery.WriteResponse(w, http.StatusBadRequest, "Incorrect parameters provided")
		return
	}
	if err == service.ErrPasswordUnableToHash {
		delivery.WriteResponse(w, http.StatusInternalServerError, "Couldn't hash password")
		return
	}
	if err == service.ErrExistingEmail {
		delivery.WriteResponse(w, http.StatusConflict, "This email is already registered")
		return
	}
	if err != nil {
		delivery.WriteResponse(w, http.StatusInternalServerError, "")
		return
	}

	response := struct {
		ID uint `json:"ID"`
	}{ID: id}
	delivery.WriteJSONResponse(w, http.StatusCreated, response)
}

func (con userControllerImpl) CheckCredentials(w http.ResponseWriter, r *http.Request) {
	var loginReq struct {
		Email    string `json:"Email"`
		Password string `json:"Password"`
	}
	err := delivery.ReadJSONRequest(r, &loginReq)
	if err != nil {
		delivery.WriteResponse(w, http.StatusBadRequest, "Incorrect request format")
		return
	}

	err = con.serv.CheckCredentials(loginReq.Email, loginReq.Password)
	if err == service.ErrIncorrectParameters {
		delivery.WriteResponse(w, http.StatusBadRequest, "Incorrect format of request")
		return
	}
	if err == service.ErrPasswordUnableToHash {
		delivery.WriteResponse(w, http.StatusInternalServerError, "Couldn't hash password")
		return
	}
	if err == service.ErrNotExistingEntity {
		delivery.WriteResponse(w, http.StatusBadRequest, "There's no user for this email")
		return
	}
	if err == service.ErrNotValidCredentials {
		delivery.WriteResponse(w, http.StatusBadRequest, "Incorrect password")
		return
	}
	if err != nil {
		delivery.WriteResponse(w, http.StatusInternalServerError, "")
		return
	}

	user, err := con.serv.GetByEmail(loginReq.Email)
	if err == service.ErrIncorrectParameters {
		delivery.WriteResponse(w, http.StatusBadRequest, "Invalid email")
		return
	}
	if err == service.ErrNotExistingEntity {
		delivery.WriteResponse(w, http.StatusNotFound, "There's no user for this email")
		return
	}
	if err != nil {
		delivery.WriteResponse(w, http.StatusInternalServerError, "")
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
	})

	tokenStr, err := token.SignedString(config.GetParams().AuthSecret)
	if err != nil {
		delivery.WriteResponse(w, http.StatusInternalServerError, "Couldn't sign token")
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "jwtToken",
		Value:    tokenStr,
		Expires:  time.Now().Add(time.Minute * 10),
		HttpOnly: true,
		Secure:   true,
	})

	delivery.WriteResponse(w, http.StatusOK, "Authenticated correctly")
}

func (con userControllerImpl) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := delivery.ParseUintParam(r, "userid")
	if err != nil {
		delivery.WriteResponse(w, http.StatusBadRequest, "Invalid id provided")
		return
	}

	user, err := con.serv.GetByID(uint(id))
	if err == service.ErrIncorrectParameters {
		delivery.WriteResponse(w, http.StatusBadRequest, "Incorrect format of request")
		return
	}
	if err == service.ErrNotExistingEntity {
		delivery.WriteResponse(w, http.StatusBadRequest, "There's no user with this ID")
		return
	}
	if err != nil {
		delivery.WriteResponse(w, http.StatusInternalServerError, "")
		return
	}

	delivery.WriteJSONResponse(w, http.StatusOK, user)
}

func (con userControllerImpl) UpdateEmail(w http.ResponseWriter, r *http.Request) {
	id, err := delivery.ParseUintParam(r, "userid")
	if err != nil {
		delivery.WriteResponse(w, http.StatusBadRequest, "Invalid id provided")
		return
	}

	var updateReq struct {
		Email string `json:"Email"`
	}
	err = delivery.ReadJSONRequest(r, &updateReq)
	if err != nil {
		delivery.WriteResponse(w, http.StatusBadRequest, "Incorrect format of request")
		return
	}

	err = con.serv.UpdateEmail(id, updateReq.Email)
	if err == service.ErrIncorrectParameters {
		delivery.WriteResponse(w, http.StatusBadRequest, "Incorrect format of request")
		return
	}
	if err == service.ErrNotExistingEntity {
		delivery.WriteResponse(w, http.StatusBadRequest, "There's no user with this ID")
		return
	}
	if err != nil {
		delivery.WriteResponse(w, http.StatusInternalServerError, "")
		return
	}

	delivery.WriteResponse(w, http.StatusOK, "Update successful")
}

func (con userControllerImpl) UpdatePassword(w http.ResponseWriter, r *http.Request) {
	id, err := delivery.ParseUintParam(r, "userid")
	if err != nil {
		delivery.WriteResponse(w, http.StatusBadRequest, "Invalid id provided")
		return
	}

	var updateReq struct {
		Password string `json:"Password"`
	}
	err = delivery.ReadJSONRequest(r, &updateReq)
	if err != nil {
		delivery.WriteResponse(w, http.StatusBadRequest, "Incorrect format of request")
		return
	}

	err = con.serv.UpdatePassword(id, updateReq.Password)
	if err == service.ErrIncorrectParameters {
		delivery.WriteResponse(w, http.StatusBadRequest, "Incorrect format of request")
		return
	}
	if err == service.ErrNotExistingEntity {
		delivery.WriteResponse(w, http.StatusBadRequest, "There's no user with this ID")
		return
	}
	if err == service.ErrPasswordUnableToHash {
		delivery.WriteResponse(w, http.StatusInternalServerError, "Couldn't hash password")
		return
	}
	if err != nil {
		delivery.WriteResponse(w, http.StatusInternalServerError, "")
		return
	}

	delivery.WriteResponse(w, http.StatusOK, "Update successful")
}

func (con userControllerImpl) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := delivery.ParseUintParam(r, "userid")
	if err != nil {
		delivery.WriteResponse(w, http.StatusBadRequest, "Invalid id provided")
		return
	}

	err = con.serv.Delete(id)
	if err == service.ErrNotExistingEntity {
		delivery.WriteResponse(w, http.StatusBadRequest, "There's no user with this ID")
		return
	}
	if err != nil {
		delivery.WriteResponse(w, http.StatusInternalServerError, "")
		return
	}

	delivery.WriteResponse(w, http.StatusOK, "Deleted")
}

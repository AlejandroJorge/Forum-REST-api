package middleware

import (
	"net/http"
	"strings"

	"github.com/AlejandroJorge/forum-rest-api/config"
	"github.com/AlejandroJorge/forum-rest-api/delivery"
	"github.com/AlejandroJorge/forum-rest-api/repository"
	"github.com/AlejandroJorge/forum-rest-api/service"
)

func Auth(next http.HandlerFunc) http.HandlerFunc {
	serv := service.NewUserService(repository.NewSQLiteUserRepository(config.SQLiteDatabase()))
	return func(w http.ResponseWriter, r *http.Request) {
		tokenStr := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
		if tokenStr == "" {
			delivery.WriteResponse(w, http.StatusBadRequest, "No auth header provided")
			return
		}

		id, err := delivery.ParseUintParam(r, "userid")
		if err != nil {
			delivery.WriteResponse(w, http.StatusBadRequest, "Invalid user ID provided")
			return
		}

		err = serv.Authorize(id, tokenStr)
		if err == service.ErrNotExistingEntity {
			delivery.WriteResponse(w, http.StatusNotFound, "The user doesn't exist")
			return
		}
		if err == service.ErrNotValidCredentials {
			delivery.WriteResponse(w, http.StatusUnauthorized, "You're not authorized to this resource")
			return
		}
		if err != nil {
			delivery.WriteResponse(w, http.StatusInternalServerError, "")
			return
		}

		next(w, r)
	}
}

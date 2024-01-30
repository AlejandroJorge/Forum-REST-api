package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/AlejandroJorge/forum-rest-api/util"
	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
)

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authCookie, err := r.Cookie("jwtToken")
		if err != nil {
			util.WriteResponse(w, http.StatusBadRequest, "No auth cookie provided")
			return
		}

		tokenStr := authCookie.Value

		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("Not corresponding signing method")
			}

			return []byte(os.Getenv("AUTH_SECRET")), nil
		})
		if err != nil {
			fmt.Println(err)
			util.WriteResponse(w, http.StatusBadRequest, "Invalid authentication token")
			return
		}

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

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			util.WriteResponse(w, http.StatusBadRequest, "Invalid claims")
			return
		}

		rawIssuerID, ok := claims["iss"]
		if !ok {
			util.WriteResponse(w, http.StatusBadRequest, "Invalid claims")
			return
		}

		issuerID := uint(rawIssuerID.(float64))

		userID := uint(id)

		if issuerID == userID {
			next(w, r)
		} else {
			util.WriteResponse(w, http.StatusUnauthorized, "You're not authorized for this resource")
		}
	}
}

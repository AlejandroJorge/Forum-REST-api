package middleware

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/AlejandroJorge/forum-rest-api/config"
	"github.com/AlejandroJorge/forum-rest-api/delivery"
	"github.com/golang-jwt/jwt"
)

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authCookie, err := r.Cookie("jwtToken")
		if err != nil {
			delivery.WriteResponse(w, http.StatusBadRequest, "No auth cookie provided")
			return
		}

		tokenStr := authCookie.Value

		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("Not corresponding signing method")
			}

			return config.GetParams().AuthSecret, nil
		})
		if err != nil {
			fmt.Println(err)
			delivery.WriteResponse(w, http.StatusBadRequest, "Invalid authentication token")
			return
		}

		id, err := delivery.ParseUintParam(r, "id")
		if err != nil {
			delivery.WriteResponse(w, http.StatusBadRequest, "Invalid ID")
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			delivery.WriteResponse(w, http.StatusBadRequest, "Invalid claims")
			return
		}

		rawIssuerID, ok := claims["iss"]
		if !ok {
			delivery.WriteResponse(w, http.StatusBadRequest, "Invalid claims")
			return
		}

		issuerID := rawIssuerID.(float64)

		if uint(issuerID) == id {
			next(w, r)
		} else {
			delivery.WriteResponse(w, http.StatusUnauthorized, "You're not authorized for this resource")
		}
	}
}

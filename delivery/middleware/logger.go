package middleware

import (
	"net/http"

	"github.com/AlejandroJorge/forum-rest-api/logging"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logging.LogRequest(r)
		next.ServeHTTP(w, r)
	})
}

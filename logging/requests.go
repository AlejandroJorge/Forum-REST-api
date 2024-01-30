package logging

import (
	"io"
	"log"
	"net/http"
)

func LogRequest(r *http.Request) {
	msg := `
  [REQUEST] Route: %s,
  [REQUEST] Method: %s,
  [REQUEST] Content: %s,
  `
	content, err := io.ReadAll(r.Body)
	if err != nil {
		content = []byte("ERROR, couldn't read body")
	}

	log.Printf(msg, r.URL.Path, r.Method, content)
}

func LogResponse(statusCode int, data interface{}) {
	msg := `
  [RESPONSE] Status code: %d
  [RESPONSE] Content: 
  %s
  `
	log.Printf(msg, statusCode, data)
}

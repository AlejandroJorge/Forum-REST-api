package logging

import (
	"log"
	"net/http"
)

func LogRequest(r *http.Request) {
	msg := `
  [REQUEST] Route: %s,
  [REQUEST] Method: %s,
  `
	log.Printf(msg, r.URL.Path, r.Method)
}

func LogResponse(statusCode int, data interface{}) {
	msg := `
  [RESPONSE] Status code: %d
  [RESPONSE] Content: 
  %s
  `
	log.Printf(msg, statusCode, data)
}

func LogRawResponse(statusCode int, message string) {
	msg := `
  [RESPONSE] Status code: %d
  [RESPONSE] Content: %s
  `
	log.Printf(msg, statusCode, message)
}

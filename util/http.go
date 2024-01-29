package util

import (
	"encoding/json"
	"net/http"
)

func ReadJSONRequest(r *http.Request, data interface{}) error {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(data)
	if err != nil {
		return ErrIncorrectParameters
	}

	return nil
}

func WriteJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	encoder := json.NewEncoder(w)
	w.WriteHeader(statusCode)
	err := encoder.Encode(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Unable to write response"))
	}
}

func WriteResponse(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	w.Write([]byte(message))
}

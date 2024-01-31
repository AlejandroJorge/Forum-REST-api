package delivery

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/AlejandroJorge/forum-rest-api/logging"
	"github.com/gorilla/mux"
)

func ReadJSONRequest(r *http.Request, data interface{}) error {
	decoder := json.NewDecoder(r.Body)
	return decoder.Decode(data)
}

func WriteJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	encoder := json.NewEncoder(w)
	w.WriteHeader(statusCode)
	err := encoder.Encode(data)
	if err != nil {
		statusCode = http.StatusInternalServerError
		w.WriteHeader(statusCode)
		errorMsg := "Unable to write response"
		w.Write([]byte(errorMsg))
		logging.LogRawResponse(statusCode, errorMsg)
	}
	logging.LogResponse(statusCode, data)
}

func WriteResponse(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	w.Write([]byte(message))
	logging.LogRawResponse(statusCode, message)
}

func ParseUintParam(r *http.Request, key string) (uint, error) {
	params := mux.Vars(r)
	valueStr, ok := params[key]
	if !ok {
		return 0, errors.New("Parameter doesn't exist")
	}

	value, err := strconv.ParseUint(valueStr, 10, 64)
	if err != nil {
		return 0, err
	}

	return uint(value), nil
}

func ParseStringParam(r *http.Request, key string) (string, error) {
	params := mux.Vars(r)
	value, ok := params[key]
	if !ok {
		return "", errors.New("Parameter doesn't exist")
	}

	return value, nil
}

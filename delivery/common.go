package delivery

import (
	"encoding/json"
	"net/http"

	"github.com/AlejandroJorge/forum-rest-api/logging"
	"github.com/AlejandroJorge/forum-rest-api/util"
)

func ReadJSONRequest(r *http.Request, data interface{}) error {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(data)
	if err != nil {
		return util.ErrIncorrectParameters
	}

	return nil
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
		logging.LogResponse(statusCode, errorMsg)
	}
}

func WriteResponse(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	w.Write([]byte(message))
	logging.LogResponse(statusCode, message)
}

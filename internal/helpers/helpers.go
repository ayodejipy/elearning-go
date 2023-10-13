package helpers

import (
	"encoding/json"
	"log"
	"net/http"
)

type BaseResponse struct {
	Success bool `json:"success"`
	Message string `json:"message"`
	Data interface{} `json:"data,omitempty"`
}

func catch(msg string, err error) {
    if err != nil {
        log.Fatalf("%v: %v", msg, err)
    }
}

// respondWithJSON responds with JSON
func RespondWithJSON(w http.ResponseWriter, code int, payload BaseResponse) {
	// encodes payload into JSON
	response, err := json.Marshal(payload)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	// set headers and write response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func RespondWithError(w http.ResponseWriter, code int, msg string) {
	RespondWithJSON(w, code, BaseResponse{ Success: false, Message: msg})
}
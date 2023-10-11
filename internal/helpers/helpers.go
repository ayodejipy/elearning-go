package helpers

import (
	"encoding/json"
	"log"
	"net/http"
)

// type BaseResponse struct {
// 	Success bool `json:"success"`
// 	Data interface{} `json:"data"`
// }

// type SuccessResponse struct {
// 	BaseResponse
// }

// type ErrResponse struct {
// 	Message string `json:"message"`
// 	BaseResponse
// }

func catch(msg string, err error) {
    if err != nil {
        log.Fatalf("%v: %v", msg, err)
    }
}

// respondWithJSON responds with JSON
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
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
	RespondWithJSON(w, code, map[string]string{"message": msg})
}
package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

// ToJson return a JSON response
func ToJson(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatal(err)
	}
}

// Error return an error response
func Error(w http.ResponseWriter, statusCode int, err error) {
	ToJson(w, statusCode, struct {
		Error string `json:"error"`
	}{
		Error: err.Error(),
	})
}

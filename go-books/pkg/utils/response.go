package utils

import (
	"encoding/json"
	"net/http"
)

func ResponseJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(map[string]string{"message": "success"})
	w.Header().Add("Content-Type", "apllication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

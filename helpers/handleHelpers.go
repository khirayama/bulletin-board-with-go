package helpers

import (
	"encoding/json"
	"net/http"
)

type Model interface {
}

func BuildJSON(w http.ResponseWriter, data Model) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

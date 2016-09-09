package main

import (
	"encoding/json"
	"net/http"
)

type payload interface {
}

func BuildJSON(w http.ResponseWriter, payload payload) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(payload)
}

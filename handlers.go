package main

import (
	"./models"
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"net/http"
)

func MessageIndex(w http.ResponseWriter, r *http.Request) {
	var messages models.Messages

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	messages = models.GetAllMessages()
	json.NewEncoder(w).Encode(messages)
}

func MessageCreate(w http.ResponseWriter, r *http.Request) {
	var message models.Message
	body, _ := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	json.Unmarshal(body, &message)
	models.CreateMessage(message)
}

func MessageDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["messageId"]
	models.DeleteMessage(id)
}

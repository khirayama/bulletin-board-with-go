package handlers

import (
	"../models/message"
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"net/http"
)

func MessageIndex(w http.ResponseWriter, r *http.Request) {
	var messages message.Messages

	messages = message.All()
	BuildJSON(w, messages)
}

func MessageCreate(w http.ResponseWriter, r *http.Request) {
	var message_ message.Message
	body, _ := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	json.Unmarshal(body, &message_)
	message_.Save()
	// message.New(message_.Name, message_.Text)
}

func MessageDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	message_ := message.Find(id)
	message_.Destroy()
	// message.Delete(id)
}

func BuildJSON(w http.ResponseWriter, data message.Messages) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

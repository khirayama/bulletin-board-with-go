package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

func MessagesIndex(w http.ResponseWriter, r *http.Request) {
	var messages Messages

	messages = getMessages()
	BuildJSON(w, messages)
}

func MessageCreate(w http.ResponseWriter, r *http.Request) {
	var message_ Message
	body, _ := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	json.Unmarshal(body, &message_)
	createMessage(message_)
}

func MessageEdit(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	message_ := findMessage(id)
	BuildJSON(w, message_)
}

func MessageUpdate(w http.ResponseWriter, r *http.Request) {
	var message_ Message
	vars := mux.Vars(r)
	id := vars["id"]
	body, _ := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	json.Unmarshal(body, &message_)
	message_.Id, _ = strconv.Atoi(id)
	updateMessage(message_)
}

func MessageDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	destroyMessage(id)
}

package handlers

import (
	"../models/message"
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

func MessageIndex(w http.ResponseWriter, r *http.Request) {
	var messages message.Messages

	messages = message.All()
	BuildJSONs(w, messages)
}

func MessageCreate(w http.ResponseWriter, r *http.Request) {
	var message_ message.Message
	body, _ := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	json.Unmarshal(body, &message_)
	message_.Save()
	// message.New(message_.Name, message_.Text)
}

func MessageEdit(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	message_ := message.Find(id)
	BuildJSON(w, message_)
}

func MessageUpdate(w http.ResponseWriter, r *http.Request) {
	var message_ message.Message
	vars := mux.Vars(r)
	id := vars["id"]
	body, _ := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	json.Unmarshal(body, &message_)
	// TODO: make id received by mux int
	message_.Id, _ = strconv.Atoi(id)
	message_.Update()
}

func MessageDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	message_ := message.Find(id)
	message_.Destroy()
	// message.Delete(id)
}

func BuildJSON(w http.ResponseWriter, data message.Message) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func BuildJSONs(w http.ResponseWriter, data message.Messages) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

func Index(w http.ResponseWriter, r *http.Request) {
	var messages Messages

	messages = All()
	BuildJSON(w, messages)
}

func Create(w http.ResponseWriter, r *http.Request) {
	var message_ Message
	body, _ := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	json.Unmarshal(body, &message_)
	message_.Save()

	// Can this also
	// message.New(message_.Name, message_.Text)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	message_ := Find(id)
	BuildJSON(w, message_)
}

func Update(w http.ResponseWriter, r *http.Request) {
	var message_ Message
	vars := mux.Vars(r)
	id := vars["id"]
	body, _ := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	json.Unmarshal(body, &message_)
	// FIXME: make id received by mux int
	message_.Id, _ = strconv.Atoi(id)
	message_.Update()

	// Can this also
	// message_.Save()
}

func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	message_ := Find(id)
	message_.Destroy()

	// Can this also
	// message.Delete(id)
}

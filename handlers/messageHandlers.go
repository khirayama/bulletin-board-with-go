package messageHandlers

import (
	"../helpers"
	"../models"
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

func Index(w http.ResponseWriter, r *http.Request) {
	var messages message.Messages

	messages = message.All()
	helpers.BuildJSON(w, messages)
}

func Create(w http.ResponseWriter, r *http.Request) {
	var message_ message.Message
	body, _ := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	json.Unmarshal(body, &message_)
	message_.Save()

	// Can this also
	// message.New(message_.Name, message_.Text)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	message_ := message.Find(id)
	helpers.BuildJSON(w, message_)
}

func Update(w http.ResponseWriter, r *http.Request) {
	var message_ message.Message
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
	message_ := message.Find(id)
	message_.Destroy()

	// Can this also
	// message.Delete(id)
}

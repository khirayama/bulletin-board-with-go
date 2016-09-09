package main

import (
	"fmt"
	"github.com/markbates/goth/gothic"
	"net/http"
)

var currentUser = User{}

func sessionsCreate(res http.ResponseWriter, req *http.Request) {
	user, err := gothic.CompleteUserAuth(res, req)
	if err != nil {
		panic(err)
	}

	currentUser := FindOrCreateUserFromAuthHash(user)
	http.Redirect(res, req, "/board", http.StatusFound)
	// fmt.Fprintln(res, "logged in!", currentUser)
	fmt.Println(currentUser)
	// fmt.Fprintln(res, "logged in!", user)
}

// import (
// 	"encoding/json"
// 	"github.com/gorilla/mux"
// 	"io"
// 	"io/ioutil"
// 	"net/http"
// 	"strconv"
// )
//
//
// func UserCreate(w http.ResponseWriter, r *http.Request) {
// 	var user_ User
// 	body, _ := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
//
// 	json.Unmarshal(body, &user_)
// 	createUser(user_)
// }
//
// func UserEdit(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	id := vars["id"]
// 	user_ := findUser(id)
// 	BuildJSON(w, user_)
// }
//
// func UserUpdate(w http.ResponseWriter, r *http.Request) {
// 	var user_ User
// 	vars := mux.Vars(r)
// 	id := vars["id"]
// 	body, _ := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
//
// 	json.Unmarshal(body, &user_)
// 	user_.Id, _ = strconv.Atoi(id)
// 	updateUser(user_)
// }
//
// func UserDelete(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	id := vars["id"]
// 	destroyUser(id)
// }

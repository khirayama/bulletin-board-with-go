package main

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

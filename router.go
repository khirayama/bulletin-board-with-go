package main

import (
	"github.com/gorilla/mux"
	"github.com/markbates/goth/gothic"
	"html/template"
	"net/http"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", HomeHandler).Methods("GET")
	router.HandleFunc("/auth/{provider}", gothic.BeginAuthHandler).Methods("GET")
	router.HandleFunc("/auth/{provider}/callback", sessionsCreate).Methods("GET")
	router.HandleFunc("/board", boardHandler).Methods("GET")

	http.Handle("/", router)
	return router
}

func HomeHandler(res http.ResponseWriter, req *http.Request) {
	tmpl, _ := template.ParseFiles("static/index.html")
	tmpl.Execute(res, nil)
}

func boardHandler(res http.ResponseWriter, req *http.Request) {
	tmpl, _ := template.ParseFiles("static/index.html")
	tmpl.Execute(res, nil)
}

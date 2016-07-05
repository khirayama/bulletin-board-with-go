package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/users", UserCreate).Methods("POST")
	router.HandleFunc("/users/{id}", UserEdit).Methods("GET")
	router.HandleFunc("/users/{id}", UserUpdate).Methods("PUT")
	router.HandleFunc("/users/{id}", UserDelete).Methods("DELETE")

	router.HandleFunc("/messages", MessagesIndex).Methods("GET")
	router.HandleFunc("/messages", MessageCreate).Methods("POST")
	router.HandleFunc("/messages/{id}", MessageEdit).Methods("GET")
	router.HandleFunc("/messages/{id}", MessageUpdate).Methods("PUT")
	router.HandleFunc("/messages/{id}", MessageDelete).Methods("DELETE")

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	http.Handle("/", router)
	return router
}

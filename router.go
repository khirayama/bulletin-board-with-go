package main

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/twitter"
	"html/template"
	"net/http"
	"os"
)

func NewRouter() *mux.Router {
	gothic.Store = sessions.NewFilesystemStore(os.TempDir(), []byte("goth-example"))
	goth.UseProviders(twitter.New(os.Getenv("TWITTER_KEY"), os.Getenv("TWITTER_SECRET"), "http://localhost:8080/auth/twitter/callback?provider=twitter"))

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", HomeHandler).Methods("GET")
	router.HandleFunc("/auth/{provider}", gothic.BeginAuthHandler).Methods("GET")
	router.HandleFunc("/auth/{provider}/callback", SessionsCreate).Methods("GET")
	// router.HandleFunc("/users/{id}", UserEdit).Methods("GET")
	// router.HandleFunc("/users/{id}", UserUpdate).Methods("PUT")
	// router.HandleFunc("/users/{id}", UserDelete).Methods("DELETE")

	// router.HandleFunc("/messages", MessagesIndex).Methods("GET")
	// router.HandleFunc("/messages", MessageCreate).Methods("POST")
	// router.HandleFunc("/messages/{id}", MessageEdit).Methods("GET")
	// router.HandleFunc("/messages/{id}", MessageUpdate).Methods("PUT")
	// router.HandleFunc("/messages/{id}", MessageDelete).Methods("DELETE")

	// router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	http.Handle("/", router)
	return router
}

func HomeHandler(res http.ResponseWriter, req *http.Request) {
	tmpl, _ := template.ParseFiles("static/index.html")
	tmpl.Execute(res, nil)
}

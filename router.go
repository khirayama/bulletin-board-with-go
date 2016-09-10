package main

import (
	"encoding/hex"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/markbates/goth/gothic"
	"html/template"
	"net/http"
	"os"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", HomeHandler).Methods("GET")
	router.HandleFunc("/auth/{provider}", gothic.BeginAuthHandler).Methods("GET")
	router.HandleFunc("/auth/{provider}/callback", sessionsCreate).Methods("GET")

	http.Handle("/", router)
	return router
}

func HomeHandler(res http.ResponseWriter, req *http.Request) {
	for _, cookie := range req.Cookies() {
		fmt.Println(cookie.Name, cookie.Value)
		if cookie.Name == "session" {
			userId := decript(os.Getenv("ENCRIPT_KEY"), []byte(cookie.Value))
			fmt.Println(hex.EncodeToString(userId))
		}
	}
	tmpl, _ := template.ParseFiles("static/index.html")
	tmpl.Execute(res, nil)
}

func boardHandler(res http.ResponseWriter, req *http.Request) {
	tmpl, _ := template.ParseFiles("static/index.html")
	tmpl.Execute(res, nil)
}

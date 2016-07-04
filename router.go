package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter(routing Routes) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routing {
		var handler http.Handler
		handler = route.HandlerFunc

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	http.Handle("/", router)
	return router
}

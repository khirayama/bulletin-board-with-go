package main

import (
	"log"
	"net/http"
)

func main() {
	log.Print("starting app...")
	router := NewRouter(Routing)

	log.Print("localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

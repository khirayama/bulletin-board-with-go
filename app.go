package main

import (
	"./base"
	"log"
	"net/http"
)

func main() {
	log.Print("starting app...")
	router := base.NewRouter()

	log.Print("localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

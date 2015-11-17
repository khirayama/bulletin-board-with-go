package main

import (
	"./config"
	"./helpers"
	"log"
	"net/http"
)

func main() {
	log.Print("starting app...")
	router := helpers.NewRouter(config.Routing)

	log.Print("localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

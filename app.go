package main

import (
	"./base"
	"log"
	"net/http"
)

func main() {
	router := base.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}

package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
)

var database, err = sql.Open("sqlite3", "./db/app.db")

func main() {
	log.Print("starting app...")
	router := NewRouter(Routing)

	log.Print("localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

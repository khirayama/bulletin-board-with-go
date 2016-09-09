package main

import (
	"database/sql"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/twitter"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"os"
)

var database, err = sql.Open("sqlite3", "./db/app.db")

func init() {
	goth.UseProviders(
		twitter.New(os.Getenv("TWITTER_KEY"),
			os.Getenv("TWITTER_SECRET"),
			"http://localhost:8080/auth/twitter/callback?provider=twitter"),
	)
}

func main() {
	log.Print("starting app...")
	router := NewRouter()

	log.Print("http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

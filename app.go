package main

import (
	"crypto/aes"
	"crypto/cipher"
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

func encrypt(key string, plaintext []byte) []byte {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	return ciphertext
}

func decript(key string, ciphertext []byte) []byte {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)

	return ciphertext
}

package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"github.com/markbates/goth/gothic"
	"net/http"
	"os"
)

func encrypt(key string, plaintext []byte) []byte {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]

	// encrypt
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	return ciphertext
}

func sessionsCreate(res http.ResponseWriter, req *http.Request) {
	user, err := gothic.CompleteUserAuth(res, req)
	if err != nil {
		panic(err)
	}

	currentUser := FindOrCreateUserFromAuthHash(user)

	encriptKey := os.Getenv("ENCRIPT_KEY")
	token := hex.EncodeToString(encrypt(encriptKey, make([]byte, currentUser.Id)))
	fmt.Printf("%x", token)

	http.SetCookie(res, &http.Cookie{
		Name:  "session",
		Value: token,
		Path:  "/",
	})
	http.Redirect(res, req, "/", http.StatusFound)
}

package main

import (
	"encoding/hex"
	"github.com/markbates/goth/gothic"
	"net/http"
	"os"
)

func sessionsCreate(res http.ResponseWriter, req *http.Request) {
	user, err := gothic.CompleteUserAuth(res, req)
	if err != nil {
		panic(err)
	}

	currentUser := FindOrCreateUserFromAuthHash(user)

	encriptKey := os.Getenv("ENCRIPT_KEY")
	token := hex.EncodeToString(encrypt(encriptKey, make([]byte, currentUser.Id)))

	http.SetCookie(res, &http.Cookie{
		Name:  "session",
		Value: token,
		Path:  "/",
	})
	http.Redirect(res, req, "/", http.StatusFound)
}

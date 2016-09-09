package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

var commonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}

func main() {
	userid := []byte("hirayama_user_id")
	key := "testtest"

	cipher_, _ := aes.NewCipher([]byte(key))

	// encrypt
	encrypter := cipher.NewCFBEncrypter(cipher_, commonIV)
	encryptedText := make([]byte, len(userid))
	encrypter.XORKeyStream(encryptedText, userid)
	fmt.Printf("%s => %x\n", userid, encryptedText)

	// decrypt
	decrypter := cipher.NewCFBDecrypter(cipher_, commonIV)
	decrypedText := make([]byte, len(userid))
	decrypter.XORKeyStream(decrypedText, encryptedText)
	fmt.Printf("%x => %s\n", encryptedText, decrypedText)
}

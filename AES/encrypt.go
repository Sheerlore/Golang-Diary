package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"io/ioutil"
)

func main() {
	fmt.Println("Encryption Program v0.01")
	

	text := []byte("My Super Secret code Stuff")
	key := []byte("passphrasewhichneedstobe32bytes!")

	// generate a new aes cipher using our 32 byte long key
	c, err := aes.NewCipher(key)

	//if there are any errors, handle them
	if err != nil {
		fmt.Println(err)
	}

	//gcm or Galois/Counter Mode, is a mode of operation
	//for symmetric key cryptographic block ciphers
	// - https://en.wikipedia.org/wiki/Galois/Counter_Mode
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		fmt.Println(err)
	}

	//creates a new byte array the size of the nonce
	//which must be passed to Seal
	nonce := make([]byte, gcm.NonceSize())
	//populates our nonce with a cryptographically secure
	//random sequence
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println(err)
	}

	//here we encrypt our text using the Seal function
	//Seal encrypts and authenticetes plaintext, authenticates the
	//additional data and appedns the result to dst, returning the updated
	//slice. The nonce must be NonceSize() bytes long and unique for all
	//time, for a given key.
	fmt.Println(gcm.Seal(nonce, nonce, text, nil))

	//the Writefile method returns an error if unsuccessful
	err = ioutil.WriteFile("myfile.data", gcm.Seal(nonce, nonce, text, nil), 0777)
	if err != nil {
		fmt.Println(err)
	}

}
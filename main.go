package main

import (
	_ "bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

func main() {
	a := newAES256WithRandomKey()
	fmt.Println(a)
}

// Create new aes with random keyb
func newAES256WithRandomKey() cipher.Block {
	key := make([]byte, 32)

	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		panic(err)
	}

	// use cipher
	cip, err := aes.NewCipher(key)

	if err != nil {
		panic(err)
	}

	return cip
}

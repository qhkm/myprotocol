package myprotocol

import (
	_ "bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

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

func padPKC7(b []byte) error {
	padSize := aes.BlockSize - (len(b) % aes.BlockSize)
	pad := bytes.Repeat([]byte{byte(padSize)}, padSize)
	return append(b, pad...)
}

func unpadPKC7(b []byte) []byte {
	padSize := int(b[len(b)-1])
	miaw := b[:len(b)-padS]
	fmt.Println(miaw)
}

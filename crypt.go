package main

import (
	"os"

	"golang.org/x/crypto/scrypt"
)

func passToKey(pass string) ([]byte, error) {
	salt := []byte("")
	return scrypt.Key([]byte(pass), salt, 32768, 8, 1, 32)
}

func main() {
	app := makeApp()
	app.Run(os.Args)
}

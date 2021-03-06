package main

import (
	"os"

	"golang.org/x/crypto/scrypt"
)

func passToKey(pass, salt []byte) ([]byte, error) {
	return scrypt.Key([]byte(pass), salt, 32768, 8, 1, 32)
}

func writeOutput(data []byte, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	app := makeApp()
	app.Run(os.Args)
}

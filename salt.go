package main

import (
	"crypto/rand"
	"os"
)

// Salt encapsulates the functionality related to generating salts
// can be easily casted to []byte like so []byte(Salt{})
type Salt []byte

// New creates a cryptographically strong random salt of the size len
func NewSalt(len int) (Salt, error) {
	buffer := make([]byte, len)
	_, err := rand.Read(buffer)
	if err != nil {
		return Salt{}, err
	}
	return Salt(buffer), nil
}

// Store stores the salt in a file given by the name, with 0600 perms : not sure about the perms tbh
func (s Salt) Store(filename string) error {
	return os.WriteFile(filename, []byte(s), 0600)
}

// Load loads the Salt from the given file
func LoadSalt(filename string) (Salt, error) {
	buffer, err := os.ReadFile(filename)
	if err != nil {
		return Salt{}, err
	}
	return Salt(buffer), nil
}

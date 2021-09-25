package main

import (
	"crypto/aes"
	"crypto/cipher"
)

// SaltGen is a func that generates salt on demand, only exist to make the signatures less verbose
type SaltGen func() (Salt, error)

// NewAESGCM creates a new AEAD cipher based on passphrase and calls SaltGen to generate a salt,
// also returns the salt for further processing by the calling code
func NewAESGCM(pass string, sg SaltGen) (cipher.AEAD, Salt, error) {
	salt, err := sg()
	if err != nil {
		return nil, nil, err
	}

	key, err := passToKey([]byte(pass), salt)
	if err != nil {
		return nil, nil, ErrKeyFailed
	}

	aesCipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, nil, ErrCipherFailed
	}

	gcm, err := cipher.NewGCM(aesCipher)
	if err != nil {
		return nil, nil, ErrGCMFailed
	}

	return gcm, salt, nil
}

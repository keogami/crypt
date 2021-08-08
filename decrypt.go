package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"io/ioutil"
	"os"

	cli "github.com/urfave/cli/v2"
)

func decryptMain(c *cli.Context) error {
	if c.String(passphraseOption) == "" {
		return ErrPassEmpty
	}
	if !c.Args().Present() {
		return ErrNoArgs
	}

	fname := c.Args().First()
	input, err := os.Open(fname)
	if err != nil {
		return err
	}
	defer input.Close()
	inb, err := ioutil.ReadAll(input)
	if err != nil {
		return err
	}

	pass := c.String(passphraseOption)
	key, err := passToKey(pass)
	if err != nil {
		return err
	}

	aesCipher, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	gcm, err := cipher.NewGCM(aesCipher)
	if err != nil {
		return err
	}

	nonce := inb[:gcm.NonceSize()]
	cipher := inb[gcm.NonceSize():]

	plain, err := gcm.Open(nil, nonce, cipher, nil)
	if err != nil {
		return err
	}

	fmt.Println(string(plain))
	return nil
}
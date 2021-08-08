package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"io/ioutil"
	"os"

	cli "github.com/urfave/cli/v2"
)

func encryptMain(c *cli.Context) error {
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
	plain, err := ioutil.ReadAll(input)
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

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return err
	}

	cipher := gcm.Seal(nil, nonce, plain, nil)

	output := os.Stdout

	outb := append(nonce, cipher...)
	_, err = output.Write(outb)
	if err != nil {
		return nil
	}

	return nil
}

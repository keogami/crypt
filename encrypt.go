package main

import (
	"crypto/rand"
	"fmt"
	"io"

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
	plain, err := ReadFile(fname, ExitArgLoadFailed)
	if err != nil {
		return err
	}

	pass := c.String(passphraseOption)

	gcm, salt, err := NewAESGCM(pass, func() (Salt, error) {
		s, err := NewSalt(saltSize)
		if err != nil {
			return nil, ErrSaltGenFailed
		}
		return s, nil
	})
	if err != nil {
		return err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return err
	}

	cipher := gcm.Seal(nil, nonce, plain, nil)

	outb := append(nonce, cipher...)
	path := c.Path(outputOption)

	if path == "" {
		fmt.Println(string(outb))
		return nil
	}

	if err := writeOutput(salt, c.Path(saltOption)); err != nil {
		return err
	}

	return writeOutput(outb, path)
}

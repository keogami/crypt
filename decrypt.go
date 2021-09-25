package main

import (
	"fmt"
	"io/ioutil"
	"os"

	cli "github.com/urfave/cli/v2"
)

func decryptMain(c *cli.Context) error {
	if c.String(passphraseOption) == "" {
		return ErrPassEmpty
	}
	if c.Path(saltOption) == "" {
		return ErrSaltEmpty
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

	gcm, _, err := NewAESGCM(pass, func() (Salt, error) {
		s, err := LoadSalt(c.Path(saltOption))
		if err != nil {
			return nil, ErrSaltLoadFailed
		}
		return s, nil
	})
	if err != nil {
		return err
	}

	nonce := inb[:gcm.NonceSize()]
	cipher := inb[gcm.NonceSize():]

	plain, err := gcm.Open(nil, nonce, cipher, nil)
	if err != nil {
		if err.Error() != "cipher: message authentication failed" {
			return err
		}
		return ErrAuthFailed
	}

	path := c.Path(outputOption)
	if path == "" {
		fmt.Println(string(plain))
		return nil
	}
	return writeOutput(plain, path)
}

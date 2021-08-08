package main

import (
	"fmt"
	"os"

	cli "github.com/urfave/cli/v2"
	"golang.org/x/term"
)

func inputPassphrase(ctx *cli.Context) error {
	if ctx.String(passphraseOption) != "" {
		return nil
	}

	fmt.Print("Input Passphrase: ")

	p, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		return nil
	}

	return ctx.Set("passphrase", string(p))
}

func makeApp() *cli.App {
	return &cli.App{
		Name:   "crypt",
		Usage:  "encrypt and decrypt files with AES-256",
		Before: inputPassphrase,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "passphrase",
				Usage:    "passphrase to be used for operations; if empty, you will be prompted to type it in",
				Aliases:  []string{"pass", "p"},
				Required: true,
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "encrypt",
				Aliases: []string{"en", "e"},
				Usage:   "encrypt the given file and outputs to stdout",
				Action:  encryptMain,
			},
			{
				Name:    "decrypt",
				Aliases: []string{"de", "d"},
				Usage:   "decrypt the given file and outputs to stdout (on failure outputs nothing)",
				Action:  decryptMain,
			},
		},
	}
}

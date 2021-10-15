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

	if !ctx.IsSet("output") {
		return ErrNoOutputPath
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
			&cli.PathFlag{
				Name:     "salt",
				Usage:    "the file where the salt is to be stored or loaded from",
				Aliases:  []string{"s"},
				Required: true,
			},
			&cli.PathFlag{
				Name:    "output",
				Usage:   "output is the path to the file where the output is stored; it is required if -passphrase is empty",
				Aliases: []string{"out", "o"},
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
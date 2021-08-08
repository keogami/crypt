package main

import (
	cli "github.com/urfave/cli/v2"
)

func makeApp() *cli.App {
	return &cli.App{
		Name:  "crypt",
		Usage: "encrypt and decrypt files with AES-256",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "passphrase",
				Usage:    "passphrase to be used for operations",
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

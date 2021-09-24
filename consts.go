package main

import (
	cli "github.com/urfave/cli/v2"
)

const (
	saltSize         = 4096
	passphraseOption = "passphrase"
	outputOption     = "output"
	saltOption       = "salt"
)

const (
	ExitPassEmpty = iota - 1
	ExitSaltEmpty
	ExitNoArgs
	ExitNoOutputPath
	ExitAuthFailed
)

var (
	ErrPassEmpty    = cli.Exit("password is empty", ExitPassEmpty)
	ErrSaltEmpty    = cli.Exit("salt file is not provided", ExitSaltEmpty)
	ErrNoArgs       = cli.Exit("not enough argument provided", ExitNoArgs)
	ErrNoOutputPath = cli.Exit("no output path provided with empty passphrase", ExitNoOutputPath)
	ErrAuthFailed   = cli.Exit("authentication failed", ExitAuthFailed)
)

package main

import (
	cli "github.com/urfave/cli/v2"
)

const (
	passphraseOption = "passphrase"
	outputOption     = "output"
)

const (
	ExitPassEmpty = iota + 1
	ExitNoArgs
	ExitNoOutputPath
	ExitAuthFailed
)

var (
	ErrPassEmpty    = cli.Exit("password is empty", ExitPassEmpty)
	ErrNoArgs       = cli.Exit("not enough argument provided", ExitNoArgs)
	ErrNoOutputPath = cli.Exit("no output path provided with empty passphrase", ExitNoOutputPath)
	ErrAuthFailed   = cli.Exit("authentication failed", ExitAuthFailed)
)

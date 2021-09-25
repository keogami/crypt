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
	ExitKeyFailed
	ExitCipherFailed
	ExitGCMFailed
	ExitSaltGenFailed
	ExitSaltLoadFailed
)

var (
	ErrPassEmpty      = cli.Exit("password is empty", ExitPassEmpty)
	ErrSaltEmpty      = cli.Exit("salt file is not provided", ExitSaltEmpty)
	ErrNoArgs         = cli.Exit("not enough argument provided", ExitNoArgs)
	ErrNoOutputPath   = cli.Exit("no output path provided with empty passphrase", ExitNoOutputPath)
	ErrAuthFailed     = cli.Exit("authentication failed", ExitAuthFailed)
	ErrKeyFailed      = cli.Exit("key derivation failed", ExitKeyFailed)
	ErrCipherFailed   = cli.Exit("aes cipher creation failed", ExitCipherFailed)
	ErrGCMFailed      = cli.Exit("GCM creation failed", ExitGCMFailed)
	ErrSaltGenFailed  = cli.Exit("salt generation failed", ExitSaltGenFailed)
	ErrSaltLoadFailed = cli.Exit("salt loading failed", ExitSaltLoadFailed)
)

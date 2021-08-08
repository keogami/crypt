package main

import (
  cli "github.com/urfave/cli/v2"
)

const passphraseOption = "passphrase"

const (
  ExitPassEmpty = iota + 1
  ExitNoArgs
)

var (
  ErrPassEmpty = cli.Exit("password is empty", ExitPassEmpty)
  ErrNoArgs = cli.Exit("not enough argument provided", ExitNoArgs)
)
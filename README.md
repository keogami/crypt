# crypt
A simple utility to encrypt &amp; decrypt files with AES-256 with passphrases

## Features
- Encrypt files using a single passphrase
- Decryption requires the passphrase &amp; a randomly-generated cryptographically-strong salt
- Easy to wrap my two brain cells around

## Installation
First install `go` from [https://go.dev/dl](https://go.dev/dl) then run the following commands:
```bash
go install github.com/keogami/crypt/cmd/crypt@latest
```
That's the end of it :3

## Running crypt
Run crypt by itself to get usage info:
```bash
crypt
```

```
NAME:
   crypt - encrypt and decrypt files with AES-256

USAGE:
   crypt [global options] command [command options] [arguments...]

COMMANDS:
   encrypt, en, e  encrypt the given file and outputs to stdout
   decrypt, de, d  decrypt the given file and outputs to stdout (on failure outputs nothing)
   help, h         Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --passphrase value, --pass value, -p value  passphrase to be used for operations; if empty, you will be prompted to type it in
   --salt value, -s value                      the file where the salt is to be stored or loaded from
   --output value, --out value, -o value       output is the path to the file where the output is stored, defaults to stdout
   --help, -h                                  show help (default: false)
```

## Encrypting a file
To encrypt files you need to specify three things:
1. A passphrase using `-p`
2. An output path using `-o`
3. A path where you want the salt to be stored using `-s`

```
crypt -p <your-passphrase> -o <output-path> -s <salt-path> encrypt <input-path>
```

For example:
```
crypt -p "my super secure passphrase" -o output.enc -s output.salt encrypt mysecretfile.txt
```

This will generate two new files:
1. `output.enc`, the encrypted file
2. `output.salt`, this file is required to decrypt `output.enc`

Now you can safely delete the original file ;3

NOTE: To avoid compromising your passphrase, omit `-p` or supply an empty string to `-p` and `crypt` will prompt you to enter your password.
```
crypt -p "" -o output.enc -s output.salt encrypt mysecretfile.txt
```

## Decrypting a file

To decrypt, you simply do the inverse of encryption:
```
crypt -p <your-passphrase> -o <decrypted-output-path> -s <salt-path> decrypt <encrypted-input-file>
```
Now you should have your original file back ;3

----

# What now?
Now you can hide stuff without much hassle X3 thank me later

# go-file-encryption

## Setup
1) Download the codebase
2) `go mod tidy` to install dependencies
3) Usage: `go run . help` for commands available

## Important function
io.ReadAll - To read all the contents of the file.
io.ReadFull - To read the contents of the file till the size given.

## Cryptographic flow
```
Read the files -> create a random nonce -> password to encrypt/decrypt -> crypto algo(SHA-1) -> iterations -> Password-Based Key Derivation Function -> AES Cipher -> Galois Counter Mode -> aesgcm(SEAL/OPEN)
```
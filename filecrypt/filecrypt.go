package filecrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha1"
	"fmt"
	"io"
	"os"

	"golang.org/x/crypto/pbkdf2"
)

func Encrypt(source string, password []byte) {
	if _, err := os.Stat(source); os.IsNotExist(err) {
		panic(err.Error())
	}

	srcFile, err := os.Open(source)

	if err != nil {
		panic(err.Error())
	}

	defer srcFile.Close()

	plainTxt, err := io.ReadAll(srcFile)

	if err != nil {
		panic(err.Error())
	}

	key := password

	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	// paswword based key derived function - key, nonce, iterations, size, crypto algo
	dk := pbkdf2.Key(key, nonce, 4096, 32, sha1.New)

	block, err := aes.NewCipher(dk)
	if err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)

	if err != nil {
		panic(err.Error())
	}

	// dst, nonce, plainText, additonaldata
	cipherText := aesgcm.Seal(nil, nonce, plainTxt, nil)

	cipherText = append(cipherText, nonce...)

	dstFile, err := os.Create(source)

	if err != nil {
		panic(err.Error())
	}

	defer dstFile.Close()

	_, err = dstFile.Write(cipherText)

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("___ENCRYPTION_DONE___")
}

func Decrypt(file []byte, password []byte) {

}

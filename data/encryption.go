package data

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"log"
	"os"

	"github.com/lunarisnia/pman/config"
)

func dirname() string {
	path, err := os.Getwd()
	if err != nil {
		log.Fatalf("get current dir err: %v", err.Error())
	}

	return path
}

func readKey() []byte {
	// TODO: This return the current path where the command is called not where the program is, change it!
	key, err := os.ReadFile(dirname() + "/pman-key.txt")
	if err != nil {
		log.Fatalf("keyfile read err: %v", err.Error())
	}
	return key
}

func createBlock() cipher.Block {
	key := readKey()
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatalf("cipher err: %v", err.Error())
	}
	return block
}

func generateNonce(nonceSize int) []byte {
	nonce := make([]byte, nonceSize)

	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		log.Fatalf("nonce err: %v", err.Error())
	}

	return nonce
}

func readFile(path string) []byte {
	plainText, _ := os.ReadFile(path)

	return plainText
}

func EncryptFile() {
	block := createBlock()
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Fatalf("cipher GCM err: %v", err.Error())
	}

	nonce := generateNonce(gcm.NonceSize())
	plainText := readFile(config.DBPATH)

	cipherText := gcm.Seal(nonce, nonce, plainText, nil)
	err = os.WriteFile(config.DBPATH, cipherText, 0777)
	if err != nil {
		log.Fatalf("write file err: %v", err.Error())
	}
}

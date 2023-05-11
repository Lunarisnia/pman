package data

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha1"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/lunarisnia/pman/config"
	"golang.org/x/crypto/pbkdf2"
)

func dirname() string {
	path, err := os.Executable()
	if err != nil {
		log.Fatalf("get current dir err: %v", err.Error())
	}

	return filepath.Dir(path)
}

func readKey() []byte {
	key, err := os.ReadFile(dirname() + "/pman-key.txt")
	hashedKey := pbkdf2.Key(key, key, 4096, 32, sha1.New)
	if err != nil {
		log.Fatalf("keyfile read err: %v", err.Error())
	}
	return hashedKey
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

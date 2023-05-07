package data

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"log"
	"os"
)

func dirname() string {
	path, err := os.Getwd()
	if err != nil {
		log.Fatalf("get current dir err: %v", err.Error())
	}

	return path
}

func readKey() []byte {
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

func deleteFile() {
	emptyBytes := []byte{}
	// TODO: change this to the new path system
	if err := os.WriteFile(dirname()+"/pman-vault.db", emptyBytes, 0777); err != nil {
		log.Fatalf("delete file err: %v", err.Error())
	}
}

func EncryptFile() {
	block := createBlock()
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Fatalf("cipher GCM err: %v", err.Error())
	}

	nonce := generateNonce(gcm.NonceSize())
	plainText := readFile(dirname() + "/pman-vault.db")

	cipherText := gcm.Seal(nonce, nonce, plainText, nil)
	err = os.WriteFile(dirname()+"/pman-vault.pman", cipherText, 0777)
	if err != nil {
		log.Fatalf("write file err: %v", err.Error())
	}

	deleteFile()
}

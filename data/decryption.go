package data

import (
	"crypto/cipher"
	"log"
	"os"
)

func getNonce(cipherText []byte, nonceSize int) []byte {
	return cipherText[:nonceSize]
}

func DecryptFile() {
	cipherText := readFile(dirname() + "/pman-vault.pman")
	block := createBlock()

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Fatalf("cipher GCM err: %v", err.Error())
	}

	nonce := getNonce(cipherText, gcm.NonceSize())
	cipherText = cipherText[gcm.NonceSize():]

	plainText, err := gcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		log.Fatalf("decrypt file err: %v", err.Error())
	}

	err = os.WriteFile(dirname()+"/pman-vault.db", plainText, 0777)
	if err != nil {
		log.Fatalf("write file err: %v", err.Error())
	}
}

package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
)

//AESEncrypt function
func AESEncrypt(stringToEncrypt string, keyString string) (encryptedString string) {
	key, _ := hex.DecodeString(keyString)
	plaintext := []byte(stringToEncrypt)
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)
	return fmt.Sprintf("%x", ciphertext)
}

//AESDecrypt function
func AESDecrypt(encryptedString string, keyString string) (decryptedString string) {
	key, _ := hex.DecodeString(keyString)
	enc, _ := hex.DecodeString(encryptedString)
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonceSize := aesGCM.NonceSize()
	nonce, ciphertext := enc[:nonceSize], enc[nonceSize:]
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		fmt.Println("Wrong Passphrase!")
	}
	return fmt.Sprintf("%s", plaintext)
}

//StringToSHA256 - get SHA256 hash
func StringToSHA256(str string) string {
	//strbyte := md5.Sum([]byte(str))
	strbyte := sha256.Sum256([]byte(str))
	hash := hex.EncodeToString(strbyte[:])
	return hash
}

package serverutil

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
)

//EncodeBase64 encapsulates the encoding from the bas64 library
func EncodeBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

//DecodeBase64 encapsulates the decoding from the base64 library
func DecodeBase64(s string) []byte {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return data
}

// Encrypt with specified key
func Encrypt(key, text []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	b := EncodeBase64(text)
	ciphertext := make([]byte, aes.BlockSize+len(b))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b))
	return ciphertext
}

// Decrypt with specified key
func Decrypt(key, text []byte) string {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	if len(text) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := text[:aes.BlockSize]
	text = text[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(text, text)
	return string(DecodeBase64(string(text)))
}

//RandomKey geneartes a random key of the size specified
func RandomKey(KeySize int) []byte {

	b := make([]byte, KeySize)

	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}

	return b
}

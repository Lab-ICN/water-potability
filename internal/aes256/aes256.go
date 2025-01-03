package aes256

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"io"
)

func Encrypt(plaintext string, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}
	cipher := gcm.Seal(nonce, nonce, []byte(plaintext), nil)
	return hex.EncodeToString(cipher), nil
}

func Decrypt(ciphertext string, key, iv []byte) (string, error) {
	decoded, err := hex.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	decrypted, err := gcm.Open(nil, iv, decoded, nil)
	if err != nil {
		return "", err
	}
	return string(decrypted), nil
}

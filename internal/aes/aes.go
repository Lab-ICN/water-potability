package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"
)

func Decrypt(_cipher, key, nonce []byte) (string, error) {
	if len(key) != aes.BlockSize {
		return "", errors.New("invalid aes key length")
	}
	if len(nonce) != 12 {
		return "", errors.New("invalid aes nonce length")
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	plaintext, err := aesgcm.Open(nil, nonce, _cipher, nil)
	if err != nil {
		return "", err
	}
	return string(plaintext), nil
}

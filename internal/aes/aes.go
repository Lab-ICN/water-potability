package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"
)

func DecryptAESGCM(ciphertext []byte) (string, error) {
	key := []byte("PR0T0T1P3_W4TER_")
	iv := []byte("labicnlabicn")

	if len(key) != 16 || len(iv) != 12 {
		return "", errors.New("invalid key or iv length")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	plaintext, err := aesgcm.Open(nil, iv, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

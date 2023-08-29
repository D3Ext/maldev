package crypto

// https://github.com/alinz/crypto.go/blob/main/chacha20.go

import (
	"crypto/rand"
	"errors"

	"golang.org/x/crypto/chacha20poly1305"
)

// Encrypt data using given key
func Chacha20Encrypt(data []byte, key []byte) ([]byte, error) {
	if len(key) != 32 {
		return nil, errors.New("bad key length, expected 32 bytes")
	}

	aead, err := chacha20poly1305.NewX(key)
	if err != nil {
		return nil, err
	}

	nonceSize := aead.NonceSize()
	// Select a random nonce, and leave capacity for the ciphertext.
	nonce := make([]byte, nonceSize, nonceSize+len(data)+aead.Overhead())

	_, err = rand.Read(nonce)
	if err != nil {
		return nil, err
	}

	// Encrypt the message and append the ciphertext to the nonce.
	return aead.Seal(nonce, nonce, data, nil), nil
}

// Decrypt data using given key
func Chacha20Decrypt(data []byte, key []byte) ([]byte, error) {
	aead, err := chacha20poly1305.NewX(key)
	if err != nil {
		return nil, err
	}

	nonceSize := aead.NonceSize()

	if len(data) < nonceSize {
		return nil, errors.New("ciphertext too short")
	}

	// Split nonce and ciphertext.
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]

	// Decrypt the message and check it wasn't tampered with.
	plaintext, err := aead.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		if err.Error() == "chacha20poly1305: message authentication failed" {
			return nil, errors.New("wrong key")
		}

		return nil, err
	}

	return plaintext, nil
}

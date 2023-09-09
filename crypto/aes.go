package crypto

/*

References:
https://gist.github.com/fracasula/38aa1a4e7481f9cedfa78a0cdd5f1865
https://golangdocs.com/aes-encryption-decryption-in-golang

*/

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
)

func AESEncrypt(plaintext []byte, iv []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return []byte(""), err
	}

	if string(plaintext) == "" {
		return []byte(""), errors.New("string to encrypt is empty")
	}

	ecb := cipher.NewCBCEncrypter(block, iv)
	content := []byte(plaintext)
	content = PKCS5Padding(content, block.BlockSize())
	crypted := make([]byte, len(content))
	ecb.CryptBlocks(crypted, content)

	return crypted, nil
}

func AESDecrypt(ciphertext []byte, iv []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return []byte(""), err
	}

	if len(ciphertext) == 0 {
		return []byte(""), errors.New("ciphertext cannot be empty")
	}

	ecb := cipher.NewCBCDecrypter(block, iv)
	decrypted := make([]byte, len(ciphertext))
	ecb.CryptBlocks(decrypted, ciphertext)

	return PKCS5Trimming(decrypted), nil
}

func GenerateIV() ([]byte, error) {
	iv := make([]byte, aes.BlockSize)
	_, err := rand.Read(iv)
	if err != nil {
		return []byte(""), err
	}

	return iv, nil
}

package crypto

import (
	"crypto/rc4"
)

func Rc4Encrypt(plaintext []byte, psk []byte) ([]byte, error) {
	r, err := rc4.NewCipher(psk)
	if err != nil {
		return []byte(""), err
	}

	dst := make([]byte, len(plaintext))
	r.XORKeyStream(dst, plaintext)
	return dst, nil
}

func Rc4Decrypt(ciphertext []byte, psk []byte) ([]byte, error) {
	r, err := rc4.NewCipher(psk)
	if err != nil {
		return []byte(""), err
	}
	src := make([]byte, len(ciphertext))
	r.XORKeyStream(src, []byte(ciphertext))
	return src, nil
}

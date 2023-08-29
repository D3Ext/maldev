package crypto

import (
	"encoding/hex"
	ecies "github.com/ecies/go/v2"
)

func EllipticCurveEncrypt(priv_key []byte, plaintext []byte) ([]byte, error) {
	hex_priv_key := hex.EncodeToString(priv_key)       // Encode to hex because is the right format that library accept
	k, err := ecies.NewPrivateKeyFromHex(hex_priv_key) // Generate key
	if err != nil {
		return []byte(""), err
	}

	ciphertext, err := ecies.Encrypt(k.PublicKey, plaintext) // Encrypt plaintext
	if err != nil {
		return []byte(""), err
	}

	return ciphertext, nil // Return ciphertext
}

func EllipticCurveDecrypt(priv_key []byte, ciphertext []byte) ([]byte, error) {
	hex_priv_key := hex.EncodeToString(priv_key)
	k, err := ecies.NewPrivateKeyFromHex(hex_priv_key)
	if err != nil {
		return []byte(""), err
	}

	dec, err := ecies.Decrypt(k, ciphertext)
	if err != nil {
		return []byte(""), err
	}

	return dec, nil
}

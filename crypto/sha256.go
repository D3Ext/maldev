package crypto

import (
	"crypto/sha256"
	"encoding/hex"
)

func Sha256(plaintext []byte) string {
	hash := sha256.Sum256(plaintext)
	return hex.EncodeToString(hash[:])
}

func VerifySha256(hash string, password string) bool {
	passwd_hash := Sha256([]byte(password))
	if hash == passwd_hash {
		return true
	} else {
		return false
	}
}

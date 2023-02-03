package crypto

import (
  "crypto/sha512"
  "encoding/hex"
)

func Sha512(plaintext []byte) (string) {
  hash := sha512.Sum512(plaintext)
  return hex.EncodeToString(hash[:])
}

func VerifySha512(hash string, password string) (bool) {
  passwd_hash := Sha512([]byte(password))
  if hash == passwd_hash {
    return true
  } else {
    return false
  }
}



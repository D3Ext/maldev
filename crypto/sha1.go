package crypto

import (
  "crypto/sha1"
  "encoding/hex"
)

func Sha1Hash(plaintext []byte) (string) {
  hash := sha1.Sum(plaintext)
  return hex.EncodeToString(hash[:])
}

func VerifySha1(hash string, password string) (bool) {
  passwd_hash := Sha1Hash([]byte(password))
  if hash == passwd_hash {
    return true
  } else {
    return false
  }
}


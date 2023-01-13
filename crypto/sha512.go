package crypto

import (
  "crypto/sha512"
  "encoding/hex"
)

func Sha512Hash(plaintext []byte) (string) {
  hash := sha512.Sum512(plaintext)
  return hex.EncodeToString(hash[:])
}



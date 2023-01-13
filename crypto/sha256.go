package crypto

import (
  "crypto/sha256"
  "encoding/hex"
)

func Sha256Hash(plaintext []byte) (string) {
  hash := sha256.Sum256(plaintext)
  return hex.EncodeToString(hash[:])
}


package crypto

import (
  "crypto/sha1"
  "encoding/hex"
)

func Sha1Hash(plaintext []byte) (string) {
  hash := sha1.Sum(plaintext)
  return hex.EncodeToString(hash[:])
}



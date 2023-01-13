package crypto

import (
  "crypto/md5"
  "encoding/hex"
)

func Md5Hash(plaintext []byte) (string) {
  hash := md5.Sum(plaintext)
  return hex.EncodeToString(hash[:])
}


package crypto

import (
  "crypto/md5"
  "encoding/hex"
)

func Md5Hash(plaintext []byte) (string) {
  hash := md5.Sum(plaintext)
  return hex.EncodeToString(hash[:])
}

func VerifyMd5(hash string, password string) (bool) {
  passwd_hash := Md5Hash([]byte(password))
  if hash == passwd_hash {
    return true
  } else {
    return false
  }
}


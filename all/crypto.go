package all

import "github.com/D3Ext/maldev/crypto"

func AESEncrypt(plaintext []byte, iv []byte, key []byte) ([]byte, error) {
  return crypto.AESEncrypt(plaintext, iv, key)
}

func AESDecrypt(ciphertext []byte, iv []byte, key []byte) ([]byte, error) {
  return crypto.AESDecrypt(ciphertext, iv, key)
}

func GenerateIV() ([]byte, error) {
  return crypto.GenerateIV()
}

func B32E(plaintext string) (string) {
  return crypto.B32E(plaintext)
}

func B32D(ciphertext string) (string) {
  return B32D(ciphertext)
}

func B64E(plaintext string) (string) {
  return crypto.B64E(plaintext)
}

func B64D(ciphertext string) (string) {
  return crypto.B64D(ciphertext)
}

func Bcrypt(plaintext []byte) ([]byte, error) {
  return crypto.Bcrypt(plaintext)
}

func VerifyBcrypt(hash []byte, plaintext []byte) (bool) {
  return crypto.VerifyBcrypt(hash, plaintext)
}

func EllipticCurveEncrypt(priv_key []byte, plaintext []byte) ([]byte, error) {
  return crypto.EllipticCurveEncrypt(priv_key, plaintext)
}

func EllipticCurveDecrypt(priv_key []byte, ciphertext []byte) ([]byte, error) {
  return crypto.EllipticCurveDecrypt(priv_key, ciphertext)
}

func EncryptFile(file string, psk string, iv []byte) (error) {
  return crypto.EncryptFile(file, psk, iv)
}

func DecryptFile(file string, psk string, iv []byte) (error) {
  return crypto.DecryptFile(file, psk, iv)
}

func Md5(plaintext []byte) (string) {
  return crypto.Md5(plaintext)
}

func VerifyMd5(hash string, password string) (bool) {
  return crypto.VerifyMd5(hash, password)
}

func Morse(input string) (string, error) {
  return crypto.Morse(input)
}

func Rc4Encrypt(plaintext []byte, psk []byte) ([]byte, error) {
  return crypto.Rc4Encrypt(plaintext, psk)
}

func Rc4Decrypt(ciphertext []byte, psk []byte) ([]byte, error) {
  return crypto.Rc4Decrypt(ciphertext, psk)
}

func Rot13(input string) string {
  return crypto.Rot13(input)
}

func Rot47(input string) string {
  return crypto.Rot47(input)
}

func Sha1(plaintext []byte) (string) {
  return crypto.Sha1(plaintext)
}

func VerifySha1(hash string, password string) (bool) {
  return crypto.VerifySha1(hash, password)
}

func Sha256(plaintext []byte) (string) {
  return crypto.Sha256(plaintext)
}

func VerifySha256(hash string, password string) (bool) {
  return crypto.VerifySha256(hash, password)
}

func Sha512(plaintext []byte) (string) {
  return crypto.Sha512(plaintext)
}

func VerifySha512(hash string, password string) (bool) {
  return crypto.VerifySha512(hash, password)
}

func Xor(buf []byte, xorchar byte) ([]byte) {
  return crypto.Xor(buf, xorchar)
}






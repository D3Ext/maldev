package crypto

import (
	"io/ioutil"
	"os"
)

func EncryptFile(file string, psk string, iv []byte) error {
	plaintext, err := ioutil.ReadFile(file) // Read file content as bytes
	if err != nil {
		return err
	}

	enc_bytes, err := AESEncrypt(plaintext, iv, []byte(psk)) // Encrypt bytes with AES function
	if err != nil {
		return err
	}

	f, err := os.OpenFile(file, os.O_RDWR, os.ModePerm)
	if err != nil {
		return err
	}

	_, err = f.WriteString(string(enc_bytes)) // Write ciphertext to file (overwrite)
	if err != nil {
		return err
	}

	return nil
}

func DecryptFile(file string, psk string, iv []byte) error {
	plaintext, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	dec_bytes, err := AESDecrypt(plaintext, iv, []byte(psk))
	if err != nil {
		return err
	}

	err = os.Truncate(file, 0) // Empty file to not overwrite data
	if err != nil {
		return err
	}

	f, err := os.OpenFile(file, os.O_RDWR, os.ModePerm)
	if err != nil {
		return err
	}

	_, err = f.WriteString(string(dec_bytes))
	if err != nil {
		return err
	}

	return nil
}

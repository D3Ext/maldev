package crypto

import (
	"encoding/base32"
	"log"
)

func B32E(plaintext string) string {
	return base32.StdEncoding.EncodeToString([]byte(plaintext))
}

func B32D(ciphertext string) string {
	dec, err := base32.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		log.Fatal(err)
	}

	return string(dec)
}

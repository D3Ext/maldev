package crypto

import (
	"encoding/base64"
	"log"
)

func B64E(plaintext string) string {
	return base64.StdEncoding.EncodeToString([]byte(plaintext))
}

func B64D(ciphertext string) string {
	dec, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		log.Fatal(err)
	}
	return string(dec)
}

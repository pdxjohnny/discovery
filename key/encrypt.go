package key

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
  "log"
)

func Encrypt(publicKey *rsa.PublicKey, message []byte) ([]byte, error) {
	result, err := rsa.EncryptOAEP(
		sha256.New(),
		rand.Reader,
		publicKey,
		message,
		noLabel,
	)
	if err != nil {
		log.Println("ERROR: key.Encrypt in rsa.EncryptOAEP: ", err)
		return nil, err
	}

	log.Printf("OAEP encrypted [%s]\n", result)
	log.Println()
	return result, nil
}

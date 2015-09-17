package key

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
  "log"
)

func Decrypt(privateKey *rsa.PrivateKey, message []byte) ([]byte, error) {
	result, err := rsa.DecryptOAEP(
		sha256.New(),
		rand.Reader,
		privateKey,
		message,
		noLabel,
	)
	if err != nil {
		log.Println("ERROR: key.Decrypt in rsa.DecryptOAEP: ", err)
		return nil, err
	}

	log.Printf("OAEP decrypted [%s]\n", result)
	log.Println()
	return result, nil
}

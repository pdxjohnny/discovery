package discovery

import (
	"crypto/rsa"
	"log"

	"github.com/pdxjohnny/key/decrypt"
	"github.com/pdxjohnny/key/encrypt"
	"github.com/pdxjohnny/key/load"
)

type Crypto interface {
	Key(interface{}) interface{}
}

func LoadKey(crypto Crypto, keyFile, keyType string) error {
	var key interface{}
	var err error
	switch keyType {
	case "public":
		key, err = load.LoadPublic(keyFile)
	case "private":
		key, err = load.LoadPrivate(keyFile)
	}
	if err != nil {
		log.Println("ERROR: Crypto.LoadKey loading private key: ", err)
		return err
	}
	crypto.Key(key)
	return nil
}

func Decrypt(crypto Crypto, message []byte) ([]byte, error) {
	message, err := decrypt.Decrypt(
		crypto.Key(nil).(*rsa.PrivateKey),
		message,
	)
	if err != nil {
		log.Println("ERROR: Crypto.Decrypt decrypting: ", err)
		return nil, err
	}
	return message, nil
}

func Encrypt(crypto Crypto, message []byte) ([]byte, error) {
	message, err := encrypt.Encrypt(
		crypto.Key(nil).(*rsa.PublicKey),
		message,
	)
	if err != nil {
		log.Println("ERROR: Crypto.Encrypt encrypting: ", err)
		return nil, err
	}
	return message, nil
}

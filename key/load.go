package key

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"log"
)

func LoadPemBlock(fileName string) ([]byte, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Println("ERROR: key.LoadPemBlock reading file: ", err)
		return nil, err
	}
	block, _ := pem.Decode(data)
	blockBytes, err := x509.DecryptPEMBlock(block, password)
	if err != nil {
		log.Println("ERROR: key.LoadPemBlock DecryptPEMBlock: ", err)
		return nil, err
	}
	return blockBytes, nil
}

func LoadPrivate(fileName string) (*rsa.PrivateKey, error) {
	key, err := LoadPemBlock(fileName)
	if err != nil {
		log.Println("ERROR: key.LoadPrivate LoadPemBlock: ", err)
		return nil, err
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(key)
	if err != nil {
		log.Println("ERROR: key.LoadPrivate ParsePKCS1PrivateKey: ", err)
		return nil, err
	}
	return privateKey, nil
}

func LoadPublic(fileName string) (*rsa.PublicKey, error) {
	key, err := LoadPemBlock(fileName)
	if err != nil {
		log.Println("ERROR: key.LoadPublic LoadPemBlock: ", err)
		return nil, err
	}
	publicKey, err := x509.ParsePKIXPublicKey(key)
	if err != nil {
		log.Println("ERROR: key.LoadPublic ParsePKIXPublicKey: ", err)
		return nil, err
	}
	return publicKey.(*rsa.PublicKey), nil
}

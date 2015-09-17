package discovery

import (
	"crypto/rsa"
	"log"

	"github.com/pdxjohnny/key/decrypt"
	"github.com/pdxjohnny/key/load"
)

type CryptService struct {
	ServerKey *rsa.PrivateKey
}

func NewCryptService() *CryptService {
	return &CryptService{
		ServerKey: nil,
	}
}

func (service *CryptService) LoadKey(privateKeyFile string) error {
	privateKey, err := load.LoadPrivate(privateKeyFile)
	if err != nil {
		log.Println("ERROR: CryptService.LoadKey loading private key: ", err)
		return err
	}
	service.ServerKey = privateKey
	return nil
}

func (service *CryptService) Decrypt(message []byte) ([]byte, error) {
	message, err := decrypt.Decrypt(
		service.ServerKey,
		message,
	)
	if err != nil {
		log.Println("ERROR: CryptService.Decrypt decrypting: ", err)
		return nil, err
	}
	return message, nil
}

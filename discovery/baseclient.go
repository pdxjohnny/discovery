package discovery

import (
	"crypto/rsa"
	"log"
	"net"
	"time"

	"github.com/pdxjohnny/key/encrypt"
	"github.com/pdxjohnny/key/load"
)

type BaseClient struct {
	interval  int
	conn      *net.Conn
	ServerKey *rsa.PublicKey
}

func (client *BaseClient) LoadKey(publicKeyFile string) error {
	publicKey, err := load.LoadPublic(publicKeyFile)
	if err != nil {
		log.Println("ERROR: BaseClient.LoadKey loading public key: ", err)
		return err
	}
	client.ServerKey = publicKey
	return nil
}

func (client *BaseClient) Dial(addr, port string) error {
	conn, err := net.Dial("udp", addr+":"+port)
	if err != nil {
		log.Println("ERROR: BaseClient.Dial: ", err)
		return err
	}
	client.conn = &conn
	return nil
}

func (client *BaseClient) Encrypt(message []byte) ([]byte, error) {
	message, err := encrypt.Encrypt(
		client.ServerKey,
		message,
	)
	if err != nil {
		log.Println("ERROR: BaseClient.Encrypt encrypting: ", err)
		return nil, err
	}
	return message, nil
}

func (client *BaseClient) Write(p []byte) (int, error) {
	var message = p
	if client.conn == nil {
		return 0, &NotYetDialed{}
	}
	if client.ServerKey != nil {
		p, err := client.Encrypt(p)
		if err != nil {
			log.Println("ERROR: BaseClient.Send while encrypting: ", err)
			return 0, err
		}
		message = p
	}
	_, err := (*client.conn).Write(message)
	if err != nil {
		log.Println("ERROR: BaseClient.Send while writing: ", err)
		return 0, err
	}
	return len(message), nil
}

func (client *BaseClient) Send(addr, port string, message []byte) {
	if client.conn == nil {
		err := client.Dial(addr, port)
		if err != nil {
			log.Println("ERROR: BaseClient.Send while dialing: ", err)
			return
		}
	}
	_, err := client.Write(message)
	if err != nil {
		log.Println("ERROR: BaseClient.Send while writing: ", err)
		return
	}
}

func (client *BaseClient) Online(addr, port string, message []byte) {
	interval := time.Duration(client.interval)
	for {
		client.Send(addr, port, message)
		time.Sleep(interval * time.Second)
	}
}

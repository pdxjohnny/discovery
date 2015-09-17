package discovery

import (
	"crypto/rsa"
	"log"
	"net"
	"time"

	"github.com/pdxjohnny/key/crypto"
)

type BaseClient struct {
	interval  int
	conn      *net.Conn
	serverKey *rsa.PublicKey
}

func NewBaseClient() *BaseClient {
	return &BaseClient{
		interval:  5,
		conn:      nil,
		serverKey: nil,
	}
}

func (client *BaseClient) Key(serverKey interface{}) interface{} {
	if serverKey != nil {
		client.serverKey = serverKey.(*rsa.PublicKey)
	}
	return client.serverKey
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

func (client *BaseClient) Write(p []byte) (int, error) {
	var message = p
	if client.conn == nil {
		return 0, &NotYetDialed{}
	}
	if client.serverKey != nil {
		p, err := crypto.Encrypt(client, p)
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

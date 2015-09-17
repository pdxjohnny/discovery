package discovery

import (
	"crypto/rsa"
	"log"
	"net"
)

type BaseService struct {
	port      string
	serverKey *rsa.PrivateKey
}

func NewBaseService() *BaseService {
	return &BaseService{
		port:      "0",
		serverKey: nil,
	}
}

func (service *BaseService) Key(serverKey interface{}) interface{} {
	if serverKey != nil {
		service.serverKey = serverKey.(*rsa.PrivateKey)
	}
	return service.serverKey
}

func (service *BaseService) SetPort(port string) {
	service.port = port
	log.Println("LOG: My port is now: ", port)
}

func (service *BaseService) BuffSize() int {
	return 1024
}

func (service *BaseService) Handle(buf []byte, received int, addr *net.UDPAddr) {
	log.Printf("Recvied %d %x\n", len(buf), buf)
	message, err := Decrypt(service, buf)
	if err != nil {
		log.Println("ERROR: BaseService.Handle decypting: ", err)
		return
	}
	log.Println(string(message), addr)
}

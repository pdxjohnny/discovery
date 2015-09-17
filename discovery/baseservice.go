package discovery

import (
	"log"
	"net"
)

type BaseService struct {
	*CryptService
	port string
}

func NewBaseService() *BaseService {
	return &BaseService{
		CryptService: NewCryptService(),
	}
}

func (service *BaseService) SetPort(port string) {
	service.port = port
	log.Println("LOG: My port is now: ", port)
}

func (service *BaseService) BuffSize() int {
	return 1024
}

func (service *BaseService) Handle(buf []byte, received int, addr *net.UDPAddr) {
	log.Println(string(buf), addr)
}

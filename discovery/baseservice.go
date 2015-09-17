package discovery

import (
	"log"
	"net"
)

type BaseService struct {
	port string
}

func (service *BaseService) SetPort(port string) {
	service.port = port
	log.Println("My port is now", port)
}

func (service *BaseService) BuffSize() int {
	return 1024
}

func (service *BaseService) Handle(buf []byte, received int, addr *net.UDPAddr) {
	log.Println(string(buf), addr)
}

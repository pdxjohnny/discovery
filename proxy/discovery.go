package proxy

import (
	"log"
	"net"

	"github.com/pdxjohnny/discovery/discovery"
)

type DiscoveryService struct {
	*discovery.CryptService
	proxy    Manager
	port     string
	interval int
}

func NewDiscoveryService(proxy Manager) *DiscoveryService {
	return &DiscoveryService{
		CryptService: discovery.NewCryptService(),
		proxy:       proxy,
	}
}

func (service *DiscoveryService) SetPort(port string) {
	service.port = port
	log.Println("LOG: My port is now: ", port)
}

func (service *DiscoveryService) BuffSize() int {
	return 1024
}

func (service *DiscoveryService) Handle(buf []byte, received int, addr *net.UDPAddr) {
	log.Println(string(buf), addr)
}

package proxy

import (
	"crypto/rsa"
	"log"
	"net"

	"github.com/pdxjohnny/discovery/discovery"
)

type DiscoveryService struct {
	proxy     Manager
	port      string
	interval  int
	serverKey *rsa.PrivateKey
}

func NewDiscoveryService(proxy Manager) *DiscoveryService {
	return &DiscoveryService{
		proxy: proxy,
	}
}

func (service *DiscoveryService) Key(serverKey interface{}) interface{} {
	if serverKey != nil {
		service.serverKey = serverKey.(*rsa.PrivateKey)
	}
	return service.serverKey
}

func (service *DiscoveryService) SetPort(port string) {
	service.port = port
	log.Println("LOG: My port is now: ", port)
}

func (service *DiscoveryService) BuffSize() int {
	return 1024
}

func (service *DiscoveryService) Handle(buf []byte, received int, addr *net.UDPAddr) {
	message, err := discovery.Decrypt(service, buf)
	if err != nil {
		log.Println("ERROR: DiscoveryService.Handle decypting: ", err)
		return
	}
	log.Println(string(message), addr)
}

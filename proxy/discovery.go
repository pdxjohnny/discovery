package proxy

import (
	"crypto/rsa"
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/pdxjohnny/discovery/discovery"
)

type DiscoveryService struct {
	proxy     Manager
	port      string
	password  string
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
	pass_and_port := strings.Split(string(buf), ":")
	password := pass_and_port[0]
	port := pass_and_port[1]
	if service.password == password {
		url := fmt.Sprintf(
			"http://%s:%s/",
			discovery.Host(addr),
			port,
		)
		service.proxy.Add(url)
	}
}

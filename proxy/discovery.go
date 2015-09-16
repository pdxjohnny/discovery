package proxy

import (
	"log"
	"net"
)

type ProxyDiscoveryService struct {
	proxy    ProxyManager
	port     string
	interval int
}

func NewProxyDiscoveryService(proxy ProxyManager) *ProxyDiscoveryService {
	return &ProxyDiscoveryService{
		proxy: proxy,
	}
}

func (service *ProxyDiscoveryService) Online(addr, port string, message []byte) {
}

func (service *ProxyDiscoveryService) Send(addr, port string, message []byte) {
}

func (service *ProxyDiscoveryService) SetPort(port string) {
	service.port = port
	log.Println("ProxyDiscoveryService port is now", port)
}

func (service *ProxyDiscoveryService) BuffSize() int {
	return 1024
}

func (service *ProxyDiscoveryService) Handle(buf []byte, received int, addr *net.UDPAddr) {
	log.Println(string(buf), addr)
}

package discovery

import (
	"log"
	"net"
	"time"
)

type BaseService struct {
	port string
	interval int
	conn *net.Conn
}

func (service *BaseService) Online(addr, port string, message []byte) {
	interval := time.Duration(service.interval)
	for {
		service.Send(addr, port, message)
		time.Sleep(interval * time.Second)
	}
}

func (service *BaseService) Send(addr, port string, message []byte) {
	if service.conn == nil {
		conn, err := net.Dial("udp", addr+":"+port)
		if err != nil {
			log.Println("ERROR BaseService.Send while dialing", err)
			return
		}
		service.conn = &conn
	}
	_, err := (*service.conn).Write(message)
	if err != nil {
		log.Println("ERROR BaseService.Send while writing", err)
		return
	}
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

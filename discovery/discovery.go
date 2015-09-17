package discovery

import (
	"bytes"
	"log"
	"net"
	"strings"
)

type Service interface {
	BuffSize() int
	SetPort(string)
	Handle([]byte, int, *net.UDPAddr)
}

func GetPort(conn *net.UDPConn) string {
	address := conn.LocalAddr().String()
	splitAddress := strings.Split(address, ":")
	port := splitAddress[len(splitAddress)-1]
	return port
}

func Listen(service Service, addr, port string) {
	ServerAddr, err := net.ResolveUDPAddr("udp", addr+":"+port)
	if err != nil {
		log.Println("ERROR: discovery.Listen ResolveUDPAddr: ", err)
		return
	}

	/* Now listen at selected port */
	ServerConn, err := net.ListenUDP("udp", ServerAddr)
	if err != nil {
		log.Println("ERROR: discovery.Listen ListenUDP: ", err)
		return
	}
	service.SetPort(GetPort(ServerConn))
	defer ServerConn.Close()

	buffSize := service.BuffSize()
	for {
		buf := make([]byte, buffSize)
		n, addr, err := ServerConn.ReadFromUDP(buf)
		if err != nil {
			log.Println("ERROR: ", err)
		}
		buf = bytes.Trim(buf, "\x00")
		service.Handle(buf, n, addr)
	}
}

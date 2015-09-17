package discovery

import (
	"bytes"
	"log"
	"net"
	"strings"

	"github.com/pdxjohnny/key/crypto"
)

type Service interface {
	crypto.Crypto
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
	buffSize := service.BuffSize()

	ServerAddr, err := net.ResolveUDPAddr("udp", addr+":"+port)
	if err != nil {
		log.Println("ERROR: discovery.Listen ResolveUDPAddr: ", err)
		return
	}

	ServerConn, err := net.ListenMulticastUDP("udp", nil, ServerAddr)
	defer ServerConn.Close()
	ServerConn.SetReadBuffer(buffSize)
	if err != nil {
		log.Println("ERROR: discovery.Listen ListenUDP: ", err)
		return
	}
	service.SetPort(GetPort(ServerConn))

	for {
		buf := make([]byte, buffSize)
		n, addr, err := ServerConn.ReadFromUDP(buf)
		if err != nil {
			log.Println("ERROR: discovery.Listen ReadFromUDP: ", err)
			continue
		}
		buf = bytes.Trim(buf, "\x00")
		buf, err = crypto.Decrypt(service, buf)
		if err != nil {
			log.Println("ERROR: discovery.Listen Decrypt: ", err)
			continue
		}
		service.Handle(buf, n, addr)
	}
}

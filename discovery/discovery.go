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

func Port(addr net.Addr) string {
	address := addr.String()
	splitAddress := strings.Split(address, ":")
	port := splitAddress[len(splitAddress)-1]
	return port
}

func Host(addr net.Addr) string {
	address := addr.String()
	splitAddress := strings.Split(address, ":")
	host := strings.Join(splitAddress[:len(splitAddress)-1], ":")
	return host
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
	service.SetPort(Port(ServerConn.LocalAddr()))

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

func Broadcast(interval int, keyFile, addr, port string, message []byte) {
	client := NewBaseClient()
	client.interval = interval
	crypto.LoadKey(
		client,
		keyFile,
		"public",
	)
	client.Online(
		addr,
		port,
		message,
	)
}

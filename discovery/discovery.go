package discovery

import (
	"log"
	"net"
	"os"
	"strings"
)

type Service interface {
	BuffSize() int
	SetPort(string)
	Handle([]byte, int, *net.UDPAddr)
}

type Client interface {
}

/* A Simple function to verify error */
func CheckError(err error) {
	if err != nil {
		log.Println("Error: ", err)
		os.Exit(1)
	}
}

func GetPort(conn *net.UDPConn) string {
	address := conn.LocalAddr().String()
	splitAddress := strings.Split(address, ":")
	port := splitAddress[len(splitAddress)-1]
	return port
}

func Listen(service Service, addr, port string) {
	ServerAddr, err := net.ResolveUDPAddr("udp", addr+":"+port)
	CheckError(err)

	/* Now listen at selected port */
	ServerConn, err := net.ListenUDP("udp", ServerAddr)
	CheckError(err)
	service.SetPort(GetPort(ServerConn))
	defer ServerConn.Close()

	buffSize := service.BuffSize()
	for {
		buf := make([]byte, buffSize)
		n, addr, err := ServerConn.ReadFromUDP(buf)
		if err != nil {
			log.Println("ERROR: ", err)
		}
		service.Handle(buf, n, addr)
	}
}

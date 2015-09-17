package discovery

import (
	"log"
	"net"
	"time"
)

type BaseClient struct {
	interval int
	conn     *net.Conn
}

func (client *BaseClient) Online(addr, port string, message []byte) {
	interval := time.Duration(client.interval)
	for {
		client.Send(addr, port, message)
		time.Sleep(interval * time.Second)
	}
}

func (client *BaseClient) Send(addr, port string, message []byte) {
	if client.conn == nil {
		conn, err := net.Dial("udp", addr+":"+port)
		if err != nil {
			log.Println("ERROR BaseClient.Send while dialing", err)
			return
		}
		client.conn = &conn
	}
	_, err := (*client.conn).Write(message)
	if err != nil {
		log.Println("ERROR BaseClient.Send while writing", err)
		return
	}
}

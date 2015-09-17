package discovery

import (
	"testing"
)

func TestListen(t *testing.T) {
	addr := "0.0.0.0"
	port := "43089"
	service := &BaseService{}
	Listen(service, addr, port)
}

func TestSend(t *testing.T) {
	addr := "0.0.0.0"
	port := "43089"
	message := "Mesage from me"
	client := BaseClient{}
	client.Send(addr, port, []byte(message))
}

func TestOnline(t *testing.T) {
	addr := "0.0.0.0"
	port := "43089"
	message := "Still online"
	client := BaseClient{}
	client.Online(addr, port, []byte(message))
}

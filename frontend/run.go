package frontend

import (
	"crypto/sha1"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/spf13/viper"

	"github.com/pdxjohnny/discovery/discovery"
)

func Run() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", sayhelloName)

	listner, err := net.Listen(
		"tcp",
		viper.GetString("addr")+":"+viper.GetString("port"),
	)
	if err != nil {
		log.Println("Listen:", err)
		return
	}

	port := discovery.Port(listner.Addr())
	password := sha1.Sum([]byte(viper.GetString("dPass")))
	message := []byte(fmt.Sprintf(
		"%x:%s",
		password,
		port,
	))
	go discovery.Broadcast(
		viper.GetInt("int"),
		viper.GetString("dKey"),
		viper.GetString("dAddr"),
		viper.GetString("dPort"),
		message,
	)

	err = http.Serve(listner, mux)
	if err != nil {
		log.Println("Serve: ", err)
		return
	}
}

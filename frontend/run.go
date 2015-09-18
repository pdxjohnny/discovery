package frontend

import (
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

	log.Println(listner.Addr().String())
	log.Println(discovery.Port(listner.Addr()))
	// message := fmt.Fprintf()
	// go discovery.Broadcast(
	// 	viper.GetInt("int"),
	// 	viper.GetString("dKey"),
	// 	viper.GetString("dAddr"),
	// 	viper.GetString("dPort"),
	// 	message,
	// )

	err = http.Serve(listner, mux)
	if err != nil {
		log.Println("Serve: ", err)
		return
	}
}

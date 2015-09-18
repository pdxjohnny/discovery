package proxy

import (
	"log"
	"net/http"

	"github.com/pdxjohnny/key/crypto"
	"github.com/spf13/viper"
)

func Run() {
	proxy := NewBaseManager()
	// If we only want to proxy to a single url
	if viper.GetString("url") != "" {
		proxy.Add(viper.GetString("url"))
	}

	// If we want to accept new frontends
	if viper.GetBool("discover") {
		discover := NewDiscoveryService(proxy)
		crypto.LoadKey(
			discover,
			viper.GetString("dKey"),
			"private",
		)
		discover.password = crypto.Sha(viper.GetString("dPass"), 10)
		go proxy.Discover(
			discover,
			viper.GetString("dAddr"),
			viper.GetString("dPort"),
		)
	}

	// Server the reverse proxy server
	mux := http.NewServeMux()
	mux.HandleFunc("/", Handler(proxy))
	err := http.ListenAndServe(":"+viper.GetString("port"), mux)
	if err != nil {
		log.Println("ERROR serving ReverseProxy:\t", err)
	}
}

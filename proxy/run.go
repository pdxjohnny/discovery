package proxy

import (
	"log"
	"net/http"

	"github.com/spf13/viper"
)

func Run() {
	proxy := NewBaseProxyManager()
	// If we only want to proxy to a single url
	if viper.GetString("url") != "" {
		proxy.Add(viper.GetString("url"))
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", Handler(proxy))
	err := http.ListenAndServe(":"+viper.GetString("port"), mux)
	if err != nil {
		log.Println("ERROR serving ReverseProxy:\t", err)
	}
}

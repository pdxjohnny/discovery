package proxy

import (
	"log"
	"net/http"
)

func Handler(proxy ProxyManager) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL)
		w.Header().Set("X-Proxy", "Golang")
		proxy.Random().ServeHTTP(w, r)
	}
}

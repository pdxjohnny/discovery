package proxy

import (
	"net/http"
)

func Handler(proxy Manager) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Proxy", "Golang")
		proxy.Random().ServeHTTP(w, r)
	}
}

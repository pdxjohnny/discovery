package frontend

import (
	"log"
	"net/http"

	"github.com/spf13/viper"
)

func Run() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", sayhelloName)
	err := http.ListenAndServe(":"+viper.GetString("port"), mux)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

package discovery

import (
	"github.com/spf13/viper"
)

func Run() {
	addr := viper.GetString("addr")
	port := viper.GetString("port")
	if viper.GetString("send") != "" {
		client := &BaseClient{}
		message := viper.GetString("send")
		client.Send(addr, port, []byte(message))
	} else if viper.GetString("online") != "" {
		client := &BaseClient{}
		message := viper.GetString("online")
		client.interval = viper.GetInt("int")
		client.Online(addr, port, []byte(message))
	} else {
		service := &BaseService{}
		Listen(service, addr, port)
	}
}

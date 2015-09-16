package discovery

import (
	"github.com/spf13/viper"
)

func Run() {
	addr := viper.GetString("addr")
	port := viper.GetString("port")
	service := &BaseService{}
	service.interval = viper.GetInt("int")
	if viper.GetString("send") != "" {
		message := viper.GetString("send")
		service.Send(addr, port, []byte(message))
	} else if viper.GetString("online") != "" {
		message := viper.GetString("online")
		service.Online(addr, port, []byte(message))
	} else {
		Listen(service, addr, port)
	}
}

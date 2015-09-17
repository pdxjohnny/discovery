package discovery

import (
	"github.com/spf13/viper"
)

func Run() {
	addr := viper.GetString("addr")
	port := viper.GetString("port")
	if viper.GetString("send") != "" ||
		viper.GetString("online") != "" {
		client := &BaseClient{}
		client.interval = viper.GetInt("int")
		if viper.GetString("key") != "" {
			LoadKey(client, viper.GetString("key"), "public")
		}
		if viper.GetString("send") != "" {
			client.Send(addr, port, []byte(viper.GetString("send")))
		} else if viper.GetString("online") != "" {
			client.Online(addr, port, []byte(viper.GetString("online")))
		}
	} else {
		service := &BaseService{}
		if viper.GetString("key") != "" {
			LoadKey(service, viper.GetString("key"), "private")
		}
		Listen(service, addr, port)
	}
}

package key

import (
	"github.com/spf13/viper"
)

func Run() {
	Gen(
		viper.GetString("key"),
		viper.GetInt("len"),
	)
	return
}

package conf

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
}

var GlobalConfig *Config

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME/.go-config-server")
	viper.AddConfigPath("/etc/go-config-server")

	if err := viper.ReadInConfig(); err != nil {
		log.Println(err)
	}

	if err := viper.Unmarshal(&GlobalConfig); err != nil {
		log.Println(err)
	}
}

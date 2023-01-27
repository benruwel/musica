package utils

import (
	"github.com/spf13/viper"
)

type Configuration struct {
	Port        string
	DatabaseURL string
}

var Config *Configuration

func InitConfig() error {
	viper.AddConfigPath("./")
	error := viper.ReadInConfig() // Find and read the config file

	Config = &Configuration{
		Port:        viper.GetString("http.port"),
		DatabaseURL: viper.GetString("db.url"),
	}

	return error
}

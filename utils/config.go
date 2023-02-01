package utils

import (
	"github.com/spf13/viper"
)

type Configuration struct {
	Host       string
	Port       string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
	DBSSLMode  string
}

var Config *Configuration

func InitConfig() error {
	viper.AddConfigPath(".") // optionally look for config in the working directory
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			viper.AutomaticEnv()

			Config = &Configuration{
				Host:       viper.GetString("HOST"),
				Port:       viper.GetString("PORT"),
				DBUser:     viper.GetString("DATABASE_USER"),
				DBPassword: viper.GetString("DATABASE_PASSWORD"),
				DBName:     viper.GetString("DATABASE_NAME"),
				DBPort:     viper.GetString("DATABASE_PORT"),
				DBSSLMode:  viper.GetString("DATABASE_SSLMODE"),
			}
			return nil
		}
		//  else , Config file was found but another error was produced

	}

	// Config file found and successfully parsed

	Config = &Configuration{
		Host:       viper.GetString("HOST"),
		Port:       viper.GetString("PORT"),
		DBUser:     viper.GetString("DATABASE_USER"),
		DBPassword: viper.GetString("DATABASE_PASSWORD"),
		DBName:     viper.GetString("DATABASE_NAME"),
		DBPort:     viper.GetString("DATABASE_PORT"),
		DBSSLMode:  viper.GetString("DATABASE_SSLMODE"),
	}
	return nil
}

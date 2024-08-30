package configs

import (
	"net/http"
	"time"

	"github.com/spf13/viper"
)

func Defaults() {
	// Set default values for configuration
	viper.SetDefault("SERVER_WRITE_TIMEOUT", 15)
	viper.SetDefault("SERVER_READ_TIMEOUT", 15)
	viper.SetDefault("SERVER_IDLE_TIMEOUT", 60)
	viper.SetDefault("KEY", "example-key")
	viper.SetDefault("ENVIROMENT", "dev")
	viper.SetDefault("PORT", 3003)
	viper.SetDefault("URL", "http://localhost:8080/")
	viper.SetDefault("NAME", "MinWeb")
	viper.SetDefault("DESCRIPTION", "Example app created with MW")
	viper.SetDefault("LOGGER_FILE", "log.txt")
}

func ServerConfig() *http.Server {
	return &http.Server{
		Addr:         ":" + viper.GetString("PORT"),
		WriteTimeout: time.Second * time.Duration(viper.GetInt("SERVER_WRITE_TIMEOUT")),
		ReadTimeout:  time.Second * time.Duration(viper.GetInt("SERVER_READ_TIMEOUT")),
		IdleTimeout:  time.Second * time.Duration(viper.GetInt("SERVER_IDLE_TIMEOUT")),
	}
}

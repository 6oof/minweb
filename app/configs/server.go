package configs

import (
	"net/http"
	"time"
)

func ServerConfig() *http.Server {
	return &http.Server{
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}
}

package config

import "time"

type ServerConfig struct {
	Addr           string
	MaxHeaderbytes int
	ReadTimeout,
	WriteTimeout time.Duration
}

func MakeServerConfig() ServerConfig {
	return ServerConfig{
		Addr:           getEnv("HTTP_SERVER_ADDRESS", "localhost:8080"),
		MaxHeaderbytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
}

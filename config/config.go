package config

import (
	"os"
	"strconv"
)

func getEnv(name, defaultValue string) string {
	if value, exists := os.LookupEnv(name); exists {
		return value
	}
	return defaultValue
}

func getEnvAsInt(name string, defaultValue int) int {
	valueStr := getEnv(name, strconv.Itoa(defaultValue))
	if v, err := strconv.Atoi(valueStr); err == nil {
		return v
	}
	return defaultValue
}

type Config struct {
	ServerConfig   ServerConfig
	PostgresConfig PostgresConfig
	CronConfig     CronConfig
}

func NewConfig() *Config {
	return &Config{
		ServerConfig:   MakeServerConfig(),
		PostgresConfig: MakePostgresConfig(),
		CronConfig:     MakeCronConfig(),
	}
}

package config

const (
	pgUser = "POSTGRES_USER"
	pgPass = "POSTGRES_PASSWORD"
	pgHost = "POSTGRES_HOST"
	pgPort = "POSTGRES_PORT"
	pgDB   = "POSTGRES_DB"
)

// PostgresConfig
type PostgresConfig struct {
	User,
	Password,
	Host,
	Port,
	DatabaseName string
}

// MakePostgresConfig
func MakePostgresConfig() PostgresConfig {
	return PostgresConfig{
		User:         getEnv(pgUser, "admin"),
		Password:     getEnv(pgPass, "admin"),
		Host:         getEnv(pgHost, "localhost"),
		Port:         getEnv(pgPort, "5432"),
		DatabaseName: getEnv(pgDB, "currency"),
	}
}

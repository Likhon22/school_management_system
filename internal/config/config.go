package config

import (
	"os"
	"sync"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

var (
	config *Config
	once   sync.Once
)

type Config struct {
	Version     string
	ServiceName string
	HttpPort    string
	DBCnf       DBConfig
	AuthCnf     AuthConfig
}

func loadConfig() {
	if err := godotenv.Load(); err != nil {
		log.Warn().Err(err).Msg("no .env file found, using system environment variables")
	}

	version := os.Getenv("VERSION")
	serviceName := os.Getenv("SERVICE_NAME")
	httpPort := os.Getenv("HTTP_PORT")

	config = &Config{
		Version:     version,
		ServiceName: serviceName,
		HttpPort:    httpPort,
		DBCnf:       LoadDBConfig(),
		AuthCnf:     LoadAuthConfig(),
	}

	validateMainConfig(config)
}

func GetConfig() *Config {
	once.Do(loadConfig)
	return config
}

func validateMainConfig(cfg *Config) {
	if cfg.Version == "" || cfg.ServiceName == "" || cfg.HttpPort == "" {
		log.Fatal().Msg("missing core service environment variables")
	}
}

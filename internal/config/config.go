package config

import (
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

var config *Config

type DBConfig struct {
	DBUrl        string
	MaxOpenConns int
	MaxIdleConns int
	MaxIdleTime  string
}
type JwtConfig struct {
	JwtSecret  string
	JwtExpires time.Duration
}
type Config struct {
	Version     string
	ServiceName string
	HttpPort    string
	DBCnf       DBConfig
	JwtCnf      JwtConfig
}

func loadConfig() {
	if err := godotenv.Load(); err != nil {
		log.Error().
			Err(err).
			Msg("error happening when loading from dotenv")
		os.Exit(1)
	}
	version := os.Getenv("VERSION")
	serviceName := os.Getenv("SERVICE_NAME")
	httpPort := os.Getenv("HTTP_PORT")
	dbUrl := os.Getenv("DB_URL")
	dbMaxOpenConnsStr := os.Getenv("DB_MAX_OPEN_CONNS")
	dbMaxIdleConnsStr := os.Getenv("DB_MAX_IDLE_CONNS")
	dbMaxIdleTime := os.Getenv("DB_MAX_IDLE_TIME")
	dbMaxOpenConns, err := strconv.Atoi(dbMaxOpenConnsStr)
	if err != nil {
		log.Error().Msg("Invalid DB_MAX_OPEN_CONNS value; must be an integer")
		os.Exit(1)
	}
	dbMaxIdleConns, err := strconv.Atoi(dbMaxIdleConnsStr)
	if err != nil {
		log.Error().Msg("Invalid DB_MAX_IDLE_CONNS value; must be an integer")
		os.Exit(1)
	}
	jwtSecret := os.Getenv("JWT_SECRET")
	jwtExpireStr := os.Getenv("JWT_EXPIRES")
	jwtExpire, err := time.ParseDuration(jwtExpireStr)
	if err != nil {
		log.Error().Msg("invalid duration")
	}
	config = &Config{
		Version:     version,
		ServiceName: serviceName,
		HttpPort:    httpPort,
		DBCnf: DBConfig{
			DBUrl:        dbUrl,
			MaxOpenConns: dbMaxOpenConns,
			MaxIdleConns: dbMaxIdleConns,
			MaxIdleTime:  dbMaxIdleTime,
		},
		JwtCnf: JwtConfig{
			JwtSecret:  jwtSecret,
			JwtExpires: jwtExpire,
		},
	}
	if config.DBCnf.DBUrl == "" {
		log.Error().
			Msg("Missing required environment variables: DB_URL")
		os.Exit(1)

	}
	if config.HttpPort == "" {
		log.Error().
			Msg("Missing required environment variables: HttpPort")
		os.Exit(1)

	}
	if config.ServiceName == "" {
		log.Error().
			Msg("Missing required environment variables:ServiceName ")
		os.Exit(1)

	}
	if config.Version == "" {
		log.Error().
			Msg("Missing required environment variables:Version ")
		os.Exit(1)

	}
	if config.JwtCnf.JwtSecret == "" || config.JwtCnf.JwtExpires == 0 {
		log.Error().
			Msg("Missing required environment variables:jwt ")
		os.Exit(1)

	}

}

func GetConfig() *Config {
	if config == nil {
		loadConfig()

	}
	return config
}

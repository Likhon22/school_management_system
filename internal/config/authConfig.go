package config

import (
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog/log"
)

type AuthConfig struct {
	JwtSecret             string
	JwtExpires            time.Duration
	ResetTokenExpDuration time.Duration
}

func LoadAuthConfig() AuthConfig {
	jwtSecret := os.Getenv("JWT_SECRET")
	jwtExpireStr := os.Getenv("JWT_EXPIRES")
	resetTokenExpDurationStr := os.Getenv("RESET_TOKEN_EXP_DURATION")

	jwtExpire, err := time.ParseDuration(jwtExpireStr)
	if err != nil {
		log.Fatal().Err(err).Msg("invalid JWT_EXPIRES format (e.g. '1h', '30m')")
	}

	resetTokenExpDuration, err := time.ParseDuration(resetTokenExpDurationStr)
	if err != nil {
		fmt.Println(resetTokenExpDuration)
		log.Fatal().Err(err).Msg("invalid RESET_TOKEN_EXP_DURATION; must be time")
	}

	if jwtSecret == "" {
		log.Fatal().Msg("JWT_SECRET is required")
	}

	return AuthConfig{
		JwtSecret:             jwtSecret,
		JwtExpires:            jwtExpire,
		ResetTokenExpDuration: resetTokenExpDuration,
	}
}

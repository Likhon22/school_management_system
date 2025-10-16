package middlewares

import (
	"school-management-system/internal/config"
)

type Middleware struct {
	IPLimiter      *ipLimiter
	AuthMiddleware AuthMiddleware
}

type AuthMiddleware struct {
	JwtConfig *config.JwtConfig
}

func NewAuthHandler(JwtConfig *config.JwtConfig) *AuthMiddleware {
	return &AuthMiddleware{
		JwtConfig: JwtConfig,
	}

}

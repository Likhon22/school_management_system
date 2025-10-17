package middlewares

import (
	"school-management-system/internal/config"
)

type Middleware struct {
	IPLimiter      *ipLimiter
	AuthMiddleware AuthMiddleware
}

type AuthMiddleware struct {
	AuthConfig *config.AuthConfig
}

func NewAuthHandler(authConfig *config.AuthConfig) *AuthMiddleware {
	return &AuthMiddleware{
		AuthConfig: authConfig,
	}

}

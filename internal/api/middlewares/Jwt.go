package middlewares

import (
	"net/http"
	"school-management-system/pkg/utils"
)

func (mw *Middleware) Jwt(jwtSecret string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			tokenStr := r.Header.Get("Authorization")
			token, err := utils.ValidateToken(tokenStr, jwtSecret)
			if err != nil || !token.Valid {
				http.Error(w, "unauthorized", http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})

	}
}

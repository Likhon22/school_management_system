package middlewares

import (
	"context"
	"net/http"
	"school-management-system/internal/api/contextkeys"
	"school-management-system/pkg/utils"
)

func (aw *AuthMiddleware) Jwt(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		tokenStr, err := r.Cookie("Bearer")
		if err != nil {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}
		token, err := utils.ValidateToken(tokenStr.Value, aw.JwtConfig.JwtSecret)

		if err != nil || !token.Valid {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		claims := token.Claims.(*utils.MyClaims)

		ctx := context.WithValue(r.Context(), contextkeys.UserKey, claims.Username)
		ctx = context.WithValue(ctx, contextkeys.RoleKey, claims.Role)
		ctx = context.WithValue(ctx, contextkeys.EmailKey, claims.Email)
		ctx = context.WithValue(ctx, contextkeys.UIdKey, claims.UID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})

}

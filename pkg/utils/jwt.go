package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type MyClaims struct {
	UID      int    `json:"uid"`
	Username string `json:"user"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

func SignedToken(userId int, email, username, role, jwt_secret string, jwt_expire time.Duration) (string, error) {
	if jwt_expire == 0 {
		jwt_expire = 5 * time.Minute

	}
	claims := MyClaims{
		UID:      userId,
		Username: username,
		Email:    email,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(jwt_expire)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(jwt_secret))
	if err != nil {
		return "", err

	}
	return signedToken, nil
}

func ValidateToken(tokenStr, jwtSecret string) (*jwt.Token, error) {

	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (any, error) {
		return ([]byte(jwtSecret)), nil
	})
	if err != nil {
		return nil, err

	}
	return token, nil
}

package utils

import (
	"net/http"
	"time"
)

func InvalidateCookie(w http.ResponseWriter, cookieName string) {
	http.SetCookie(w, &http.Cookie{
		Name:     cookieName,
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		Expires:  time.Unix(0, 0),
		MaxAge:   -1,
	})
}

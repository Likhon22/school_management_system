package middlewares

import (
	"net/http"
)

func (mw *Middleware) SecurityHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Prevent browsers from MIME-sniffing the content
		w.Header().Set("X-Content-Type-Options", "nosniff")

		// Prevent clickjacking
		w.Header().Set("X-Frame-Options", "DENY") // or SAMEORIGIN if you allow same-site iframes

		// Enable XSS protection in older browsers
		w.Header().Set("X-XSS-Protection", "1; mode=block")

		// Content Security Policy (CSP) - prevent XSS and data injection
		w.Header().Set("Content-Security-Policy", "default-src 'self'; script-src 'self'; style-src 'self'; img-src 'self' data:; object-src 'none'")

		// Referrer Policy - control what info browsers send in the Referer header
		w.Header().Set("Referrer-Policy", "strict-origin-when-cross-origin")

		// Feature Policy / Permissions Policy - control browser features
		w.Header().Set("Permissions-Policy", "geolocation=(), camera=(), microphone=()")

		// HSTS - force HTTPS (only set if your site uses HTTPS)
		// w.Header().Set("Strict-Transport-Security", "max-age=63072000; includeSubDomains; preload")

		next.ServeHTTP(w, r)

	})
}

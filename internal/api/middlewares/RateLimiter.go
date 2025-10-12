package middlewares

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

type ipLimiter struct {
	mu        sync.Mutex
	ips       map[string]*rate.Limiter
	tokenRate time.Duration // time per token
	burst     int           // max tokens at once
}

// Initialize with config
func NewIPLimiter(tokenRate time.Duration, burst int) *ipLimiter {
	return &ipLimiter{
		ips:       make(map[string]*rate.Limiter),
		burst:     burst,
		tokenRate: tokenRate,
	}
}

func (rl *ipLimiter) getLimiter(ip string) *rate.Limiter {
	rl.mu.Lock()
	defer rl.mu.Unlock()
	limiter, exists := rl.ips[ip]
	if !exists {
		limiter = rate.NewLimiter(rate.Every(rl.tokenRate), rl.burst)
		rl.ips[ip] = limiter

	}
	return limiter
}

func (rl *ipLimiter) RateLimiter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr
		limiter := rl.getLimiter(ip)
		if !limiter.Allow() {
			http.Error(w, "too many request", http.StatusTooManyRequests)
			return

		}
		fmt.Println("from rate limiting")
		next.ServeHTTP(w, r)

	})
}

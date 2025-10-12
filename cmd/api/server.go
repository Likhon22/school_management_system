package main

import (
	"fmt"
	"net/http"
	"os"
	"school-management-system/internal/api/middlewares"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from test!")
}
func main() {
	port := ":3000"
	fmt.Println("server running on port", port)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	mux := http.NewServeMux()
	//if want to use https
	// cert := "cert.pem"
	// pem := "key.pem"
	// tlsconfig := &tls.Config{
	// 	MinVersion: tls.VersionTLS12,
	// }
	mw := middlewares.Middleware{
		IPLimiter: middlewares.NewIPLimiter(time.Minute/12, 5),
	}
	wrappedMux := middlewares.ChainMiddleware(
		mux,
		mw.Logger,                // log everything including blocked requests
		mw.Cors,                  // must run early to avoid browser CORS errors
		mw.IPLimiter.RateLimiter, // block excessive requests before heavy processing
		mw.SecurityHeaders,       // set headers before sending response
		mw.Compression,           // compress last, after everything else is wrapped
		mw.SecurityHeaders)
	server := &http.Server{
		Addr:    port,
		Handler: wrappedMux,
		// TLSConfig: tlsconfig,
	}
	mux.Handle("GET /", http.HandlerFunc(test))
	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err)
	}
}

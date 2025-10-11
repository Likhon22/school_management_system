package main

import (
	"fmt"
	"net/http"
	"school-management-system/internal/api/middlewares"
)

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from test!")
}
func main() {
	port := ":3000"
	fmt.Println("server running on port", port)
	mux := http.NewServeMux()
	//if want to use https
	// cert := "cert.pem"
	// pem := "key.pem"
	// tlsconfig := &tls.Config{
	// 	MinVersion: tls.VersionTLS12,
	// }
	mw := middlewares.Middleware{}
	wrappedMux := middlewares.ChainMiddleware(mux, mw.SecurityHeaders)
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

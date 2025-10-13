package router

import (
	"net/http"
	"school-management-system/internal/api/handlers"
)

func Router() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("GET /", http.HandlerFunc(handlers.RootHandler))
	return mux
}

package students

import "net/http"

func (h *Handler) StudentRoutes(mux *http.ServeMux) {
	mux.Handle("POST /students", http.HandlerFunc(h.Create))

}

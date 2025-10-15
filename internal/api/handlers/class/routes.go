package class

import "net/http"

func (h *Handler) ClassRoutes(mux *http.ServeMux) {
	mux.Handle("POST /class", http.HandlerFunc(h.Create))
	mux.Handle("GET /class", http.HandlerFunc(h.Get))
	mux.Handle("GET /class/{id}", http.HandlerFunc(h.GetClassById))
	mux.Handle("PATCH /class/{id}", http.HandlerFunc(h.Update))
	mux.Handle("DELETE /class/{id}", http.HandlerFunc(h.Delete))
}

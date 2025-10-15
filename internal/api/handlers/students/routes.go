package students

import "net/http"

func (h *Handler) StudentRoutes(mux *http.ServeMux) {
	mux.Handle("POST /students", http.HandlerFunc(h.Create))
	mux.Handle("GET /students", http.HandlerFunc(h.Get))
	mux.Handle("GET /students/{id}", http.HandlerFunc(h.GetStudentById))
	mux.Handle("PATCH /students/{id}", http.HandlerFunc(h.Update))
	mux.Handle("DELETE /students/{id}", http.HandlerFunc(h.Delete))
}

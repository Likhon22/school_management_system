package subjects

import "net/http"

func (h *Handler) SubjectRoutes(mux *http.ServeMux) {
	mux.Handle("POST /subjects", http.HandlerFunc(h.Create))
	mux.Handle("GET /subjects", http.HandlerFunc(h.Get))
	mux.Handle("GET /subjects/{id}", http.HandlerFunc(h.GetSubjectById))
	mux.Handle("PATCH /subjects/{id}", http.HandlerFunc(h.Update))
	mux.Handle("DELETE /subjects/{id}", http.HandlerFunc(h.Delete))
}

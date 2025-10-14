package teachers

import "net/http"

func (h *Handler) TeacherRoutes(mux *http.ServeMux) {
	mux.Handle("POST /teachers", http.HandlerFunc(h.CREATE))
	mux.Handle("GET /teachers", http.HandlerFunc(h.Get))
	mux.Handle("GET /teachers/{id}", http.HandlerFunc(h.GetTeacherById))
	mux.Handle("PATCH /teachers/{id}", http.HandlerFunc(h.Update))

}

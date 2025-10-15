package teachers

import "net/http"

func (h *Handler) TeacherRoutes(mux *http.ServeMux) {
	mux.Handle("POST /teachers", http.HandlerFunc(h.Create))
	mux.Handle("GET /teachers", http.HandlerFunc(h.Get))
	mux.Handle("GET /teachers/{id}", http.HandlerFunc(h.GetTeacherById))
	mux.Handle("GET /teachers/{id}/students", http.HandlerFunc(h.GetStudentsByTeacherID))
	mux.Handle("PATCH /teachers/{id}", http.HandlerFunc(h.Update))
	mux.Handle("DELETE /teachers/{id}", http.HandlerFunc(h.Delete))

}

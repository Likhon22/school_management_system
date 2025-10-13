package teachers

import "net/http"

func (h *Handler) TeacherRoutes(mux *http.ServeMux) {
	mux.Handle("POST /teachers", http.HandlerFunc(h.CREATE))
}

package exec

import "net/http"

func (h *Handler) ExecsRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /execs", h.Get)
	mux.HandleFunc("POST /execs", h.CREATE)
	mux.HandleFunc("GET /execs/{id}", h.GetExecById)
	mux.HandleFunc("PATCH /execs/{id}", h.Update)
	mux.HandleFunc("DELETE /execs/{id}", h.Delete)
	mux.HandleFunc("POST /execs/{id}/updatepassword", h.CREATE)

	mux.HandleFunc("POST /execs/login", h.CREATE)
	mux.HandleFunc("POST /execs/logout", h.CREATE)
	mux.HandleFunc("POST /execs/forgotpassword", h.CREATE)
	mux.HandleFunc("POST /execs/resetpasssword/reset/{resetcode}", h.CREATE)

}

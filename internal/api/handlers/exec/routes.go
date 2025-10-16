package exec

import "net/http"

func (h *Handler) ExecsRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /execs", h.Get)
	mux.HandleFunc("POST /execs", h.Create)
	mux.HandleFunc("GET /execs/{id}", h.GetExecById)
	mux.HandleFunc("PATCH /execs/{id}", h.Update)
	mux.HandleFunc("DELETE /execs/{id}", h.Delete)
	mux.HandleFunc("POST /execs/{id}/updatepassword", h.Create)

	mux.HandleFunc("POST /execs/login", h.Create)
	mux.HandleFunc("POST /execs/logout", h.Create)
	mux.HandleFunc("POST /execs/forgotpassword", h.Create)
	mux.HandleFunc("POST /execs/resetpasssword/reset/{resetcode}", h.Create)

}

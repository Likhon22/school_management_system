package exec

import (
	"net/http"
	"school-management-system/internal/api/middlewares"
)

func (h *Handler) ExecsRoutes(mux *http.ServeMux, authMiddleware middlewares.MiddlewareFunc) {
	mux.Handle("GET /execs", authMiddleware(http.HandlerFunc(h.Get)))
	mux.HandleFunc("POST /execs", h.Create)
	mux.HandleFunc("GET /execs/{id}", h.GetExecById)
	mux.HandleFunc("PATCH /execs/{id}", h.Update)
	mux.HandleFunc("DELETE /execs/{id}", h.Delete)
	mux.HandleFunc("POST /execs/{id}/updatepassword", h.Create)

	mux.HandleFunc("POST /execs/login", h.Login)
	mux.HandleFunc("POST /execs/logout", h.Logout)
	mux.HandleFunc("POST /execs/forgotpassword", h.Create)
	mux.HandleFunc("POST /execs/resetpasssword/reset/{resetcode}", h.Create)

}

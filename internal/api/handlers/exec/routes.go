package exec

import (
	"net/http"
	"school-management-system/internal/api/middlewares"
	"school-management-system/internal/models"
)

func (h *Handler) ExecsRoutes(mux *http.ServeMux, amw *middlewares.AuthMiddleware) {
	mux.Handle("GET /execs", amw.Jwt(amw.RequiredRoles(string(models.RoleAdmin))(http.HandlerFunc(h.Get))))
	mux.HandleFunc("POST /execs", h.Create)
	mux.HandleFunc("GET /execs/{id}", h.GetExecById)
	mux.HandleFunc("PATCH /execs/{id}", h.Update)
	mux.HandleFunc("DELETE /execs/{id}", h.Delete)
	mux.Handle("POST /execs/{id}/updatepassword", amw.Jwt((http.HandlerFunc(h.UpdatePassword))))

	mux.HandleFunc("POST /execs/login", h.Login)
	mux.HandleFunc("POST /execs/logout", h.Logout)
	mux.HandleFunc("POST /execs/forgotpassword", h.Create)
	mux.HandleFunc("POST /execs/resetpasssword/reset/{resetcode}", h.Create)

}

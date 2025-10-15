package router

import (
	"net/http"
	"school-management-system/internal/api/handlers/root"
	"school-management-system/internal/api/handlers/students"
	"school-management-system/internal/api/handlers/teachers"
)

func SetupRoutes(teacherHandler *teachers.Handler, studentHandler *students.Handler) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("GET /{$}", http.HandlerFunc(root.RootHandler))
	//teacher routes
	teacherHandler.TeacherRoutes(mux)
	studentHandler.StudentRoutes(mux)

	return mux
}

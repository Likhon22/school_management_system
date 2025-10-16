package router

import (
	"net/http"
	"school-management-system/internal/api/handlers/class"
	"school-management-system/internal/api/handlers/exec"
	"school-management-system/internal/api/handlers/root"
	"school-management-system/internal/api/handlers/students"
	"school-management-system/internal/api/middlewares"

	"school-management-system/internal/api/handlers/teachers"
)

func SetupRoutes(teacherHandler *teachers.Handler, studentHandler *students.Handler, classHandler *class.Handler, execHandler *exec.Handler, amw *middlewares.AuthMiddleware) *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("GET /{$}", http.HandlerFunc(root.RootHandler))
	//teacher routes
	teacherHandler.TeacherRoutes(mux)

	//student routes
	studentHandler.StudentRoutes(mux)

	//subject routes
	classHandler.ClassRoutes(mux)

	// exec routes
	execHandler.ExecsRoutes(mux, amw)

	return mux
}

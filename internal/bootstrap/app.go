package bootstrap

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"school-management-system/internal/api/handlers/class"
	"school-management-system/internal/api/handlers/exec"
	"school-management-system/internal/api/handlers/students"
	"school-management-system/internal/api/handlers/teachers"
	"school-management-system/internal/api/middlewares"
	"school-management-system/internal/api/router"
	"school-management-system/internal/config"
	"school-management-system/internal/repository"
	"school-management-system/internal/service"
	"school-management-system/internal/validation"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type App struct {
	server *http.Server
}

func NewApp(cnf *config.Config, dbCon *sql.DB) *App {

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	mw := &middlewares.Middleware{
		IPLimiter: middlewares.NewIPLimiter(time.Minute/12, 5),
	}
	authMiddleware := middlewares.NewAuthHandler(&cnf.AuthCnf)
	log.Info().Msg("database connected successfully")
	validator := validation.NewValidator()
	//teacher handler
	teacherRepo := repository.NewTeacherRepo(dbCon)
	teacherService := service.NewTeacherService(teacherRepo)
	teacherHandler := teachers.NewHandler(teacherService, validator)

	//student handler
	studentRepo := repository.NewStudentRepo(dbCon)
	studentService := service.NewStudentService(studentRepo)
	studentHandler := students.NewHandler(studentService, validator)

	//class handler
	classRepo := repository.NewClassRepo(dbCon)
	classService := service.NewClassService(classRepo)
	classHandler := class.NewHandler(classService, validator)
	//exec handler
	execRepo := repository.NewExecRepo(dbCon)
	execService := service.NewExecService(execRepo, cnf.AuthCnf.JwtSecret, cnf.AuthCnf.JwtExpires)
	execHandler := exec.NewHandler(execService, validator)
	mux := router.SetupRoutes(teacherHandler, studentHandler, classHandler, execHandler, authMiddleware)

	wrappedMux := middlewares.SetupMiddlewares(mux, mw)
	server := &http.Server{
		Addr:    cnf.HttpPort,
		Handler: wrappedMux,
	}
	return &App{
		server: server,
	}
}

func (a *App) Run() {
	log.Info().Msgf("Server started on %s", a.server.Addr)
	if err := a.server.ListenAndServe(); err != nil {
		fmt.Println(err)
	}
}

package bootstrap

import (
	"fmt"
	"net/http"
	"os"
	"school-management-system/internal/api/handlers/students"
	"school-management-system/internal/api/handlers/teachers"
	"school-management-system/internal/api/middlewares"
	"school-management-system/internal/api/router"
	"school-management-system/internal/config"
	"school-management-system/internal/infra/db"
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

func NewApp() *App {
	cnf := config.GetConfig()
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	db, err := db.ConnectDB(cnf.DBCnf)
	if err != nil {
		log.Error().
			Err(err).
			Msg("db connection fail")
		os.Exit(1)

	}
	defer db.Close()
	log.Info().Msg("database connected successfully")
	validator := validation.NewValidator()
	//teacher handler
	teacherRepo := repository.NewTeacherRepo(db)
	teacherService := service.NewTeacherService(teacherRepo)
	teacherHandler := teachers.NewHandler(teacherService, validator)

	//student handler
	studentRepo := repository.NewStudentRepo(db)
	studentService := service.NewStudentService(studentRepo)
	studentHandler := students.NewHandler(studentService, validator)

	mux := router.SetupRoutes(teacherHandler, studentHandler)

	mw := &middlewares.Middleware{
		IPLimiter: middlewares.NewIPLimiter(time.Minute/12, 5),
	}

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

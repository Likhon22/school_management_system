package main

import (
	"fmt"
	"net/http"
	"os"
	"school-management-system/internal/api/handlers/teachers"
	"school-management-system/internal/api/middlewares"
	"school-management-system/internal/api/router"
	"school-management-system/internal/config"
	"school-management-system/internal/infra/db"
	"school-management-system/internal/repository"
	"school-management-system/pkg/utils"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
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
	//teacher handler
	teacherRepo := repository.NewTeacherRepo(db)
	teacherHandler := teachers.NewHandler(teacherRepo)
	mux := router.SetupRoutes(teacherHandler)

	mw := middlewares.Middleware{
		IPLimiter: middlewares.NewIPLimiter(time.Minute/12, 5),
	}

	wrappedMux := utils.ChainMiddleware(
		mux,
		mw.Logger,                // log everything including blocked requests
		mw.Cors,                  // must run early to avoid browser CORS errors
		mw.IPLimiter.RateLimiter, // block excessive requests before heavy processing
		mw.SecurityHeaders,       // set headers before sending response
		mw.Compression,           // compress last, after everything else is wrapped
	)
	server := &http.Server{
		Addr:    cnf.HttpPort,
		Handler: wrappedMux,
		// TLSConfig: tlsconfig,
	}
	fmt.Println("Server started on port", cnf.HttpPort)

	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err)
	}
}

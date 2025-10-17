package main

import (
	"os"
	"school-management-system/internal/bootstrap"
	"school-management-system/internal/config"
	"school-management-system/internal/infra/db"

	"github.com/rs/zerolog/log"
)

func main() {
	cnf := config.GetConfig()
	dbCon, err := db.ConnectDB(cnf.DBCnf)

	if err != nil {
		log.Error().
			Err(err).
			Msg("db connection fail")
		os.Exit(1)

	}
	defer func() {
		if err := dbCon.Close(); err != nil {
			log.Error().Err(err).Msg("error closing database")
		} else {
			log.Info().Msg("database connection closed")
		}

	}()

	app := bootstrap.NewApp(cnf, dbCon)
	app.Run()
}

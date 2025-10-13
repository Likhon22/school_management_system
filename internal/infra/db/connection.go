package db

import (
	"context"
	"database/sql"
	"school-management-system/internal/config"
	"time"

	_ "github.com/lib/pq"
)

var dbInstance *sql.DB

func ConnectDB(cfg config.DBConfig) (*sql.DB, error) {
	if dbInstance != nil {
		return dbInstance, nil
	}
	db, err := sql.Open("postgres", cfg.DBUrl)
	if err != nil {
		return nil, err

	}
	db.SetMaxOpenConns((cfg.MaxOpenConns))
	db.SetMaxIdleConns(cfg.MaxOpenConns)
	maxIdleTime, err := time.ParseDuration(cfg.MaxIdleTime)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}
	db.SetConnMaxIdleTime(maxIdleTime)
	dbInstance = db
	return db, nil
}

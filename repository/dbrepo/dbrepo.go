package dbrepo

import (
	"database/sql"
	"jamiuafolabi/web-service-gin/config"
	"jamiuafolabi/web-service-gin/repository"
)

// Holds the database connection and application configuration needed throughout the application
type DBRepo struct {
	DB  *sql.DB
	App *config.AppConfig
}

// creates a new Database repository instance
func NewDBRepo(conn *sql.DB, appconfig *config.AppConfig) repository.DatabaseRepo {
	return &DBRepo{
		App: appconfig,
		DB:  conn,
	}
}

package dbrepo

import (
	"database/sql"
	"jamiuafolabi/web-service-gin/config"
	"jamiuafolabi/web-service-gin/repository"
)

type postgresDBRepo struct {
	DB  *sql.DB
	App *config.AppConfig
}

func NewPostgresRepo(conn *sql.DB, appconfig *config.AppConfig) repository.DatabaseRepo {
	return &postgresDBRepo{
		App: appconfig,
		DB:  conn,
	}
}

package dbrepo

import (
	"database/sql"

	"github.com/milospp/pub-quiz/src/go-socket-app/internal/config"
	"github.com/milospp/pub-quiz/src/go-socket-app/internal/repository"
)

type postgresDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

func NewPostgresRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &postgresDBRepo{
		App: a,
		DB:  conn,
	}
}

package dbrepo

import (
	"database/sql"

	"github.com/milospp/pub-quiz/pkg/config"
	"github.com/milospp/pub-quiz/pkg/repository"
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

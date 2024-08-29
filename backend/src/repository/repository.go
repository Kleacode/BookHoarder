package repository

import (
	"back/src/usecases"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) usecases.RepositoryInterface {
	return Repository{db: db}
}

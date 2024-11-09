package repository

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Promocode interface {
}

type Reward interface {
}

type Repository struct {
	Promocode
	Reward
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}

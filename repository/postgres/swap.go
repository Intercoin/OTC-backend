package postgres

import "github.com/jmoiron/sqlx"

type SwapRepository struct {
	db *sqlx.DB
}

func NewSwapRepository(db *sqlx.DB) *SwapRepository {
	return &SwapRepository{db: db}
}

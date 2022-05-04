package postgres

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
)

type BlockRepository struct {
	db *sqlx.DB
}

func NewBlockRepository(db *sqlx.DB) *BlockRepository {
	return &BlockRepository{db: db}
}

type block struct {
	Network   Network   `db:"network"`
	Number    uint64    `db:"number"`
	UpdatedAt time.Time `db:"updated_at"`
}

func newBlock(network Network, number uint64) *block {
	return &block{
		Network: network,
		Number:  number,
	}
}

func (repo *BlockRepository) GetLastBlock(ctx context.Context, network Network) (uint64, error) {
	dto := &block{Network: network}
	query, args, err := repo.db.BindNamed(`
SELECT number,
       updated_at
FROM blocks
WHERE network = :network
LIMIT 1;`, dto)
	if err != nil {
		return 0, err
	}
	if err = repo.db.GetContext(ctx, dto, query, args...); err != nil {
		return 0, err
	}
	return dto.Number, nil
}

func (repo *BlockRepository) Update(ctx context.Context, network Network, number uint64) (uint64, error) {
	dto := newBlock(network, number)
	query, args, err := repo.db.BindNamed(`
UPDATE blocks
SET number = :number
WHERE network = :network
RETURNING updated_at;`, dto)
	if err != nil {
		return 0, err
	}
	if err := repo.db.GetContext(ctx, dto, query, args...); err != nil {
		return 0, err
	}
	return dto.Number, nil
}

package postgres

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
)

type SwapRepository struct {
	db *sqlx.DB
}

func NewSwapRepository(db *sqlx.DB) *SwapRepository {
	return &SwapRepository{db: db}
}

type Trade struct {
	ID        int    `db:"id"`
	Tx        string `db:"tx"`
	TradeHash string `db:"tradehash"`

	//-----------Paticipant #1-----------------
	Address1 string  `db:"address_1"`
	Network1 Network `db:"network_1"`

	//-----------Paticipant #2----------------
	Address2 string  `db:"address_2"`
	Network2 Network `db:"network_2"`

	CreatedAt time.Time `db:"created_at"`
}

type Lock struct {
	ID        int    `db:"id"`
	TradeHash string `db:"tradehash"`

	Address    string  `db:"address"`
	Network    Network `db:"network"`
	Asset      string  `db:"asset"`
	Amount     string  `db:"amount"`
	MaxPenalty string  `db:"max_penalty"`
	Deadline   string  `db:"deadline"`

	CreatedAt time.Time `db:"created_at"`
}

type Engage struct {
	ID        int    `db:"id"`
	TradeHash string `db:"tradehash"`

	Address   string  `db:"address"`
	Network   Network `db:"network"`
	Signature string  `db:"signature"`
}

type Claim struct {
	ID        int    `db:"id"`
	TradeHash string `db:"tradehash"`

	Address    string  `db:"address"`
	Network    Network `db:"network"`
	Signatures string  `db:"signatures"`
	Penalty    string  `db:"penalty"`
}

func (repo *SwapRepository) GetTrade(ctx context.Context, tradehash string) (*Trade, error) {
	dto := &Trade{TradeHash: tradehash}
	query, args, err := repo.db.BindNamed(`
SELECT id,
       tx,
       address_1,
       network_1,
       address_2,
       network_2,
       created_at
FROM trades
WHERE tradehash = :tradehash;`, dto)
	if err != nil {
		return nil, err
	}
	if err := repo.db.GetContext(ctx, dto, query, args...); err != nil {
		return nil, err
	}
	return dto, nil
}

func (repo *SwapRepository) CreateTrade(ctx context.Context, trade *Trade) error {
	query, args, err := repo.db.BindNamed(`
INSERT INTO trades (tx, tradehash, address_1, network_1, address_2) 
VALUES (:tx, 
		:tradehash, 
		:address_1, 
        :network_1,
        :address_2)
RETURNING id, created_at;`, trade)
	if err != nil {
		return err
	}
	return repo.db.GetContext(ctx, trade, query, args...)
}

func (repo *SwapRepository) UpdateTrade(ctx context.Context, tradeID int, network Network) error {
	dto := &Trade{ID: tradeID, Network2: network}
	query, args, err := repo.db.BindNamed(`
UPDATE trades
SET network_2 = :network_2
WHERE id = :id;`, dto)
	if err != nil {
		return err
	}
	_, err = repo.db.ExecContext(ctx, query, args...)
	return err
}

func (repo *SwapRepository) SaveLock(ctx context.Context, lock *Lock) error {
	query, args, err := repo.db.BindNamed(`
INSERT INTO locks (tradehash, address, network, asset, amount, max_penalty, deadline) 
VALUES (:tradehash, 
		:address, 
        :network, 
        :asset, 
        :amount, 
        :max_penalty, 
        :deadline)
RETURNING id, created_at;`, lock)
	if err != nil {
		return err
	}
	return repo.db.GetContext(ctx, lock, query, args...)
}

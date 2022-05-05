package postgres

import (
	"context"
	"database/sql"
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
	Tx        string `db:"tx"`
	TradeHash string `db:"tradehash"`

	Address    string         `db:"address"`
	Network    Network        `db:"network"`
	Asset      string         `db:"asset"`
	Amount     string         `db:"amount"`
	MaxPenalty sql.NullString `db:"max_penalty"`
	Deadline   string         `db:"deadline"`

	CreatedAt time.Time `db:"created_at"`
}

type Engage struct {
	ID        int    `db:"id"`
	Tx        string `db:"tx"`
	TradeHash string `db:"tradehash"`

	Address   string  `db:"address"`
	Network   Network `db:"network"`
	Signature string  `db:"signature"`
}

type Claim struct {
	ID        int    `db:"id"`
	Tx        string `db:"tx"`
	TradeHash string `db:"tradehash"`

	Address    string         `db:"address"`
	Network    Network        `db:"network"`
	Signatures []string       `db:"signatures"`
	Penalty    sql.NullString `db:"penalty"`
}

func (repo *SwapRepository) GetTrade(ctx context.Context, tradehash string) (*Trade, error) {
	dto := &Trade{TradeHash: tradehash}
	query, args, err := repo.db.BindNamed(`
SELECT id,
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
INSERT INTO trades (tradehash, address_1, network_1, address_2) 
VALUES (:tradehash, 
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
INSERT INTO locks (tx, tradehash, address, network, asset, amount, max_penalty, deadline) 
VALUES (:tx,
        :tradehash, 
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

func (repo *SwapRepository) SaveEngage(ctx context.Context, engage *Engage) error {
	query, args, err := repo.db.BindNamed(`
INSERT INTO engages (tx, tradehash, address, network, signature) 
VALUES (:tx, 
        :tradehash, 
		:address, 
        :network, 
        :signature)
RETURNING id;`, engage)
	if err != nil {
		return err
	}
	return repo.db.GetContext(ctx, engage, query, args...)
}

func (repo *SwapRepository) SaveClaim(ctx context.Context, claim *Claim) error {
	query, args, err := repo.db.BindNamed(`
INSERT INTO claims (tx, tradehash, address, network, penalty, signatures) 
VALUES (:tx, 
        :tradehash, 
		:address, 
        :network, 
        :penalty, 
        :signatures)
RETURNING id;`, claim)
	if err != nil {
		return err
	}
	return repo.db.GetContext(ctx, claim, query, args...)
}

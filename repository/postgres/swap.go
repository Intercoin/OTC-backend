package postgres

import (
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

type swap struct {
	ID        int    `db:"id"`
	TradeHash string `db:"tradehash"`
	//-----------Paticipant #1-----------------
	Address1       string         `db:"address_1"`
	Network1       Network        `db:"network_1"`
	Asset1         string         `db:"asset_1"`
	Amount1        int64          `db:"amount_1"`
	LockTradeHash1 sql.NullString `db:"lock_tradehash_1"`
	Withdraw1      sql.NullString `db:"withdraw_1"`
	//-----------Paticipant #2----------------
	Address2       string         `db:"address_2"`
	Network2       Network        `db:"network_2"`
	Asset2         string         `db:"asset_2"`
	Amount2        int64          `db:"amount_2"`
	LockTradeHash2 sql.NullString `db:"lock_tradehash_2"`
	Withdraw2      sql.NullString `db:"withdraw_2"`
	CreatedAt      time.Time      `db:"created_at"`
	UpdatedAt      time.Time      `db:"updated_at"`
}

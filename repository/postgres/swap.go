package postgres

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type SwapRepository struct {
	db *sqlx.DB
}

func NewSwapRepository(db *sqlx.DB) *SwapRepository {
	return &SwapRepository{db: db}
}

type trade struct {
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

type lock struct {
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

type engage struct {
	ID        int    `db:"id"`
	TradeHash string `db:"tradehash"`

	Address   string  `db:"address"`
	Network   Network `db:"network"`
	Signature string  `db:"signature"`
}

type claim struct {
	ID        int    `db:"id"`
	TradeHash string `db:"tradehash"`

	Address   string  `db:"address"`
	Network   Network `db:"network"`
	Signature string  `db:"signature"`
	Penalty   string  `db:"penalty"`
}

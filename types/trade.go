package types

import (
	"time"
)

type Trade struct {
	ID           int         `json:"id"`
	TradeHash    string      `json:"tradehash"`
	Participant1 Participant `json:"participant_1"`
	Participant2 Participant `json:"participant_2"`
	CreatedAt    time.Time   `db:"created_at"`
	UpdatedAt    time.Time   `db:"updated_at"`
}

type Participant struct {
	Address       string  `json:"address"`
	Network       string  `json:"network"`
	Asset         string  `json:"asset"`
	Amount        int64   `json:"amount"`
	LockTradeHash *string `json:"lock_tradehash,omitempty"`
	Withdraw      *string `json:"withdraw,omitempty"`
}

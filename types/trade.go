package types

import (
	"time"

	"otc-backend/repository/postgres"
)

type Network = postgres.Network

const (
	Ethereum = postgres.Ethereum
	BSC      = postgres.BSC
)

func NetworkToDB(network Network) postgres.Network {
	switch network {
	case Ethereum:
		return postgres.Ethereum
	default:
		return postgres.BSC
	}
}

type Trade struct {
	ID        int       `json:"id"`
	Tx        string    `json:"tx"`
	TradeHash string    `json:"tradehash"`
	Address1  string    `json:"address_1"`
	Network1  Network   `json:"network_1"`
	Address2  string    `json:"address_2"`
	Network2  Network   `json:"network_2"`
	CreatedAt time.Time `json:"created_at"`
}

type Lock struct {
	TradeHash  string    `json:"tradehash"`
	Address    string    `json:"address"`
	Network    Network   `json:"network"`
	Asset      string    `json:"asset"`
	Amount     string    `json:"amount"`
	MaxPenalty string    `json:"max_penalty"`
	Deadline   string    `json:"deadline"`
	CreatedAt  time.Time `json:"created_at"`
}

type Engage struct {
	TradeHash string  `json:"tradehash"`
	Address   string  `json:"address"`
	Network   Network `json:"network"`
	Signature string  `json:"signature"`
}

type Claim struct {
	TradeHash string  `json:"tradehash"`
	Address   string  `json:"address"`
	Network   Network `json:"network"`
	Signature string  `json:"signature"`
	Penalty   string  `json:"penalty"`
}

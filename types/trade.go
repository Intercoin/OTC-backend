package types

import (
	"time"

	"otc-backend/repository/postgres"
)

type Network = postgres.Network

const (
	Ethereum = postgres.Ethereum
	BSC      = postgres.BSC
	Unknown  = postgres.Unknown
)

func NetworkToDB(network Network) postgres.Network {
	switch network {
	case Ethereum:
		return postgres.Ethereum
	case BSC:
		return postgres.BSC
	default:
		return postgres.Unknown
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

func (t *Trade) ToDB() *postgres.Trade {
	return &postgres.Trade{
		ID:        t.ID,
		Tx:        t.Tx,
		TradeHash: t.TradeHash,
		Address1:  t.Address1,
		Network1:  t.Network1,
		Address2:  t.Address2,
		Network2:  t.Network2,
		CreatedAt: t.CreatedAt,
	}
}

func TradeFromDB(t *postgres.Trade) *Trade {
	return &Trade{
		ID:        t.ID,
		Tx:        t.Tx,
		TradeHash: t.TradeHash,
		Address1:  t.Address1,
		Network1:  t.Network1,
		Address2:  t.Address2,
		Network2:  t.Network2,
		CreatedAt: t.CreatedAt,
	}
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

func (t *Lock) ToDB() *postgres.Lock {
	return &postgres.Lock{
		TradeHash:  t.TradeHash,
		Address:    t.Address,
		Network:    t.Network,
		Asset:      t.Asset,
		Amount:     t.Amount,
		MaxPenalty: t.MaxPenalty,
		Deadline:   t.Deadline,
		CreatedAt:  t.CreatedAt,
	}
}

type Engage struct {
	TradeHash string  `json:"tradehash"`
	Address   string  `json:"address"`
	Network   Network `json:"network"`
	Signature string  `json:"signature"`
}

func (t *Engage) ToDB() *postgres.Engage {
	return &postgres.Engage{
		TradeHash: t.TradeHash,
		Address:   t.Address,
		Network:   t.Network,
		Signature: t.Signature,
	}
}

type Claim struct {
	TradeHash  string  `json:"tradehash"`
	Address    string  `json:"address"`
	Network    Network `json:"network"`
	Signatures string  `json:"signature"`
	Penalty    string  `json:"penalty"`
}

func (t *Claim) ToDB() *postgres.Claim {
	return &postgres.Claim{
		TradeHash:  t.TradeHash,
		Address:    t.Address,
		Network:    t.Network,
		Signatures: t.Signatures,
	}
}

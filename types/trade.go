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
		TradeHash: t.TradeHash,
		Address1:  t.Address1,
		Network1:  t.Network1,
		Address2:  t.Address2,
		Network2:  t.Network2,
		CreatedAt: t.CreatedAt,
	}
}

type Lock struct {
	Tx         string    `json:"tx"`
	TradeHash  string    `json:"tradehash"`
	Address    string    `json:"address"`
	Network    Network   `json:"network"`
	Asset      string    `json:"asset"`
	Amount     string    `json:"amount"`
	MaxPenalty string    `json:"max_penalty"`
	Deadline   string    `json:"deadline"`
	CreatedAt  time.Time `json:"created_at"`
}

func (l *Lock) ToDB() *postgres.Lock {
	return &postgres.Lock{
		Tx:         l.Tx,
		TradeHash:  l.TradeHash,
		Address:    l.Address,
		Network:    l.Network,
		Asset:      l.Asset,
		Amount:     l.Amount,
		MaxPenalty: postgres.NullStringFromString(l.MaxPenalty),
		Deadline:   l.Deadline,
		CreatedAt:  l.CreatedAt,
	}
}

func LockFromDB(l *postgres.Lock) *Lock {
	lock := &Lock{
		Tx:        l.Tx,
		TradeHash: l.TradeHash,
		Address:   l.Address,
		Network:   l.Network,
		Asset:     l.Asset,
		Amount:    l.Amount,
		Deadline:  l.Deadline,
		CreatedAt: l.CreatedAt,
	}

	if l.MaxPenalty.Valid {
		lock.MaxPenalty = l.MaxPenalty.String
	}

	return lock
}

type Engage struct {
	Tx        string  `json:"tx"`
	TradeHash string  `json:"tradehash"`
	Address   string  `json:"address"`
	Network   Network `json:"network"`
	Signature string  `json:"signature"`
}

func (e *Engage) ToDB() *postgres.Engage {
	return &postgres.Engage{
		Tx:        e.Tx,
		TradeHash: e.TradeHash,
		Address:   e.Address,
		Network:   e.Network,
		Signature: e.Signature,
	}
}

func EngageFromDB(e *postgres.Engage) *Engage {
	return &Engage{
		Tx:        e.Tx,
		TradeHash: e.TradeHash,
		Address:   e.Address,
		Network:   e.Network,
		Signature: e.Signature,
	}
}

type Claim struct {
	Tx         string   `json:"tx"`
	TradeHash  string   `json:"tradehash"`
	Address    string   `json:"address"`
	Network    Network  `json:"network"`
	Signatures []string `json:"signature"`
	Penalty    string   `json:"penalty"`
}

func (t *Claim) ToDB() *postgres.Claim {
	return &postgres.Claim{
		Tx:         t.Tx,
		TradeHash:  t.TradeHash,
		Address:    t.Address,
		Network:    t.Network,
		Penalty:    postgres.NullStringFromString(t.Penalty),
		Signatures: t.Signatures,
	}
}

func ClaimFromDB(c *postgres.Claim) *Claim {
	claim := &Claim{
		Tx:         c.Tx,
		TradeHash:  c.TradeHash,
		Address:    c.Address,
		Network:    c.Network,
		Signatures: c.Signatures,
	}

	if c.Penalty.Valid {
		claim.Penalty = c.Penalty.String
	}

	return claim
}

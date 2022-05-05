package types

import "errors"

var (
	ErrNoObject = errors.New("no object")
)

type GetTradeByTradehashResponse struct {
	Address1 string    `json:"address_1"`
	Address2 string    `json:"address_2"`
	Locks    []*Lock   `json:"locks"`
	Engages  []*Engage `json:"engages,omitempty"`
	Claims   []*Claim  `json:"claims,omitempty"`
}

package web3

import (
	"github.com/ethereum/go-ethereum/common"
	"otc-backend/types"
)

func (event *OTCSwapNewTransfer) ToModel() *types.Lock {
	return &types.Lock{
		TradeHash:  common.BytesToHash(event.TradeHash[:]).String(),
		Address:    event.Poster.String(),
		Asset:      event.Asset.String(),
		Amount:     event.Amount.String(),
		MaxPenalty: event.MaxPenalty.String(),
		Deadline:   event.Deadline.String(),
	}
}

func (event *OTCSwapNewTransfer) ToTradeModel(tx string, network types.Network) *types.Trade {
	return &types.Trade{
		Tx:        tx,
		TradeHash: common.BytesToHash(event.TradeHash[:]).String(),
		Address1:  event.Poster.String(),
		Network1:  network,
		Address2:  event.Receiver.String(),
	}
}

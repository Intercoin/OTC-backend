package web3

import (
	"github.com/ethereum/go-ethereum/common"
	"otc-backend/types"
)

func (event *OTCSwapNewTransfer) ToModel() *types.Lock {
	return &types.Lock{
		Tx:         event.Raw.TxHash.String(),
		TradeHash:  common.BytesToHash(event.TradeHash[:]).String(),
		Address:    event.Poster.String(),
		Asset:      event.Asset.String(),
		Amount:     event.Amount.String(),
		MaxPenalty: event.MaxPenalty.String(),
		Deadline:   event.Deadline.String(),
	}
}

func (event *OTCSwapNewTransfer) ToTradeModel(network types.Network) *types.Trade {
	return &types.Trade{
		TradeHash: common.BytesToHash(event.TradeHash[:]).String(),
		Address1:  event.Poster.String(),
		Network1:  network,
		Address2:  event.Receiver.String(),
	}
}

func (event *OTCSwapEngaged) ToModel() *types.Engage {
	return &types.Engage{
		Tx:        event.Raw.TxHash.String(),
		TradeHash: common.BytesToHash(event.TradeHash[:]).String(),
		Address:   event.Sender.String(),
		Signature: common.BytesToHash(event.Signature).String(),
	}
}

func (event *OTCSwapClaimed) ToModel() *types.Claim {
	return &types.Claim{
		Tx:        event.Raw.TxHash.String(),
		TradeHash: common.BytesToHash(event.TradeHash[:]).String(),
		Address:   event.Sender.String(),
		Penalty:   event.Penalty.String(),
		Signatures: []string{
			common.BytesToHash(event.Signatures[0]).String(),
			common.BytesToHash(event.Signatures[1]).String(),
		},
	}
}

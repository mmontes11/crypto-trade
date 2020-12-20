package model

import (
	"github.com/mmontes11/crypto-trade/cmd/subscriber/log"
	"github.com/mmontes11/crypto-trade/internal/core"
)

// TradeRepositoryI defines repository operations
type TradeRepositoryI interface {
	SaveTrade(t core.Trade) error
}

// TradeRepository implements repository operations
type TradeRepository struct{}

// NewTradeRepository creates a new repository instance
func NewTradeRepository() TradeRepositoryI {
	return &TradeRepository{}
}

// SaveTrade creates a new trade
func (r *TradeRepository) SaveTrade(t core.Trade) error {
	log.Logger.Debug(t)
	return nil
}

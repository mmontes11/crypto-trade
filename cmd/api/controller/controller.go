package controller

import (
	"database/sql"

	"github.com/mmontes11/crypto-trade/cmd/api/model"
	"github.com/mmontes11/crypto-trade/internal/core"
)

// TradeControllerI defines controller operations
type TradeControllerI interface {
	GetTrades() ([]core.Trade, error)
}

// TradeController implements controller operations
type TradeController struct {
	db         *sql.DB
	repository model.TradeRepositoryI
}

// NewTradeController creates a new controller instance
func NewTradeController(db *sql.DB, r model.TradeRepositoryI) TradeControllerI {
	return &TradeController{
		db,
		r,
	}
}

// GetTrades retrieves trades
func (tc *TradeController) GetTrades() ([]core.Trade, error) {
	return nil, nil
}

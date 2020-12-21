package controller

import (
	ctx "context"
	"database/sql"

	"github.com/mmontes11/crypto-trade/cmd/api/model"
	"github.com/mmontes11/crypto-trade/internal/core"
	"github.com/mmontes11/crypto-trade/pkg/database"
)

// TradeControllerI defines controller operations
type TradeControllerI interface {
	GetTrades(params core.TradeParams) ([]core.Trade, error)
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
func (tc *TradeController) GetTrades(params core.TradeParams) ([]core.Trade, error) {
	var trades []core.Trade
	txFn := func(ctx ctx.Context, tx *sql.Tx) error {
		result, err := tc.repository.GetTrades(ctx, tx, params)
		if err != nil {
			return err
		}

		trades = result
		return nil
	}

	err := database.Tx(ctx.Background(), tc.db, txFn)
	if err != nil {
		return nil, err
	}

	return trades, nil
}

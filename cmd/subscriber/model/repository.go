package model

import (
	ctx "context"
	"database/sql"

	"github.com/mmontes11/crypto-trade/internal/core"
)

// TradeRepositoryI defines repository operations
type TradeRepositoryI interface {
	SaveTrade(ctx ctx.Context, tx *sql.Tx, t core.Trade) error
}

// TradeRepository implements repository operations
type TradeRepository struct{}

// NewTradeRepository creates a new repository instance
func NewTradeRepository() TradeRepositoryI {
	return &TradeRepository{}
}

// SaveTrade creates a new trade
func (r *TradeRepository) SaveTrade(ctx ctx.Context, tx *sql.Tx, t core.Trade) error {
	query := "INSERT INTO trades (* EXCEPT(event_time)) VALUES(?, ?, ?, ?, ?)"
	stmt, err := tx.PrepareContext(ctx, query)
	defer stmt.Close()
	if err != nil {
		return err
	}

	params := []interface{}{
		t.Side,
		t.CryptoSize.Size,
		t.CryptoSize.Currency,
		t.Price.Amount,
		t.Price.Currency,
	}
	_, err = stmt.ExecContext(ctx, params...)
	if err != nil {
		return err
	}

	return nil
}

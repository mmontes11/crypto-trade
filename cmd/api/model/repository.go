package model

import (
	ctx "context"
	"database/sql"
	"time"

	"github.com/mmontes11/crypto-trade/internal/core"
)

// TradeRepositoryI defines repository operations
type TradeRepositoryI interface {
	GetTrades(ctx ctx.Context, tx *sql.Tx, params core.TradeParams) ([]core.Trade, error)
}

// TradeRepository implements repository operations
type TradeRepository struct{}

// NewTradeRepository creates a new repository instance
func NewTradeRepository() TradeRepositoryI {
	return &TradeRepository{}
}

// GetTrades retrieves trades
func (r *TradeRepository) GetTrades(ctx ctx.Context, tx *sql.Tx, params core.TradeParams) ([]core.Trade, error) {
	query := `
		SELECT
			toStartOfMinute(t.event_time) AS time,
			t.side,
			AVG(t.crypto_size) AS size,
			t.crypto_currency,
			AVG(t.price) AS price,
			t.price_currency
		FROM
			trades t
		GROUP BY
			time,
			t.side,
			t.crypto_currency,
			t.price_currency
		HAVING
			t.crypto_currency = ?
			AND t.price_currency = ?
		ORDER BY
			time DESC
		LIMIT ?
	`
	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	args := []interface{}{
		params.Crypto,
		params.Currency,
		params.Limit,
	}
	rows, err := stmt.QueryContext(ctx, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return decodeRows(rows)
}

func decodeRows(rows *sql.Rows) ([]core.Trade, error) {
	trades := []core.Trade{}
	for rows.Next() {
		var (
			time           time.Time
			side           string
			size           float64
			cryptoCurrency string
			price          float64
			priceCurrency  string
		)
		if err := rows.Scan(&time, &side, &size, &cryptoCurrency, &price, &priceCurrency); err != nil {
			return nil, err
		}

		trade := core.Trade{
			Time: time,
			Side: side,
			CryptoSize: core.CryptoSize{
				Size:     size,
				Currency: cryptoCurrency,
			},
			Price: core.Price{
				Amount:   price,
				Currency: priceCurrency,
			},
		}
		trades = append(trades, trade)
	}
	return trades, nil
}

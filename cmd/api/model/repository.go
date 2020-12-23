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
	query, args := getQuery(params)
	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return decodeRows(rows, params)
}

const (
	rawQuery = `
		SELECT time,
			t.side,
			t.size,
			t.price
		FROM trades t
		WHERE t.size_currency = ?
			AND t.price_currency = ?
		ORDER BY t.time
		LIMIT ?
	`
	aggByMinuteQuery = `
		SELECT toStartOfMinute(t.time) AS time,
			t.side,
			AVG(t.size) AS size,
			AVG(t.price) AS price
		FROM trades t
		WHERE t.size_currency = ?
			AND t.price_currency = ?
		GROUP BY t.time,
			t.side
		ORDER BY t.time
		LIMIT ?
	`
	aggByHourQuery = `
		SELECT t.time,
			t.side,
			avgMerge(t.avg_size_state) AS size,
			avgMerge(t.avg_price_state) AS price
		FROM trades_hourly t
		WHERE t.size_currency = ?
			AND t.price_currency = ?
		GROUP BY t.time,
			t.side
		ORDER BY t.time
		LIMIT ?
	`
)

func getQuery(params core.TradeParams) (query string, args []interface{}) {
	args = []interface{}{
		params.Crypto,
		params.Currency,
		params.Limit,
	}
	switch params.GroupBy {
	case "hour":
		query = aggByHourQuery
	case "minute":
		query = aggByMinuteQuery
	case "second":
		query = rawQuery
	default:
		query = rawQuery
	}
	return
}

func decodeRows(rows *sql.Rows, params core.TradeParams) ([]core.Trade, error) {
	trades := []core.Trade{}
	for rows.Next() {
		var (
			time  time.Time
			side  string
			size  float32
			price float32
		)
		if err := rows.Scan(&time, &side, &size, &price); err != nil {
			return nil, err
		}

		trade := core.Trade{
			Time: time,
			Side: side,
			Size: core.CurrencyAmount{
				Amount:   size,
				Currency: params.Crypto,
			},
			Price: core.CurrencyAmount{
				Amount:   price,
				Currency: params.Currency,
			},
		}
		trades = append(trades, trade)
	}
	return trades, nil
}

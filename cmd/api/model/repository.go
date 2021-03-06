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
		SELECT t.time,
			t.side,
			t.size,
			t.price
		FROM trades t
		WHERE t.size_currency = ?
			AND t.price_currency = ?
			AND t.time >= ?
			AND t.time <= ?
		ORDER BY (t.time, t.side)
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
			AND t.time >= ?
			AND t.time <= ?
		GROUP BY time,
			t.side
		ORDER BY (time, t.side)
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
			AND t.time >= ?
			AND t.time <= ?
		GROUP BY t.time,
			t.side
		ORDER BY (t.time, t.side)
		LIMIT ?
	`
)

func getQuery(params core.TradeParams) (query string, args []interface{}) {
	switch params.GroupBy {
	case "second":
		query = rawQuery
	case "minute":
		query = aggByMinuteQuery
	case "hour":
		query = aggByHourQuery
	default:
		query = rawQuery
	}
	args = []interface{}{
		params.Crypto,
		params.Currency,
		params.FromDate,
		params.ToDate,
		params.Limit,
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

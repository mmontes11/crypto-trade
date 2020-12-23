CREATE MATERIALIZED VIEW IF NOT EXISTS trades_hourly_mv TO trades_hourly AS
SELECT toStartOfHour(time) AS time,
	side,
	avgState(size) AS avg_size_state,
	size_currency,
	avgState(price) AS avg_price_state,
	price_currency
FROM trades
GROUP BY time,
	side,
	size_currency,
	price_currency
ORDER BY (time, side, size_currency, price_currency);
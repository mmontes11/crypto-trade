CREATE TABLE IF NOT EXISTS trades_hourly (
	time DateTime,
	side String,
	avg_size_state AggregateFunction(avg, Float32),
	size_currency String,
	avg_price_state AggregateFunction(avg, Float32),
	price_currency String
) ENGINE = SummingMergeTree() PARTITION BY (toYear(time))
ORDER BY (time, side, size_currency, price_currency);
CREATE TABLE IF NOT EXISTS trades (
    time DateTime DEFAULT now(),
    side String,
    size Float32,
    size_currency String,
    price Float32,
    price_currency String,
    total_price MATERIALIZED size * price
) Engine = MergeTree() PARTITION BY (toYYYYMMDD(time), side)
ORDER BY time;
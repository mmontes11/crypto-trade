CREATE TABLE trades (
    event_time DateTime DEFAULT now(),
    side String,
    crypto_size Float64,
    crypto_currency String,
    price Float64,
    price_currency String,
    total_price MATERIALIZED crypto_size * price
) Engine = MergeTree() PARTITION BY toYYYYMMDD(event_time)
ORDER BY event_time;
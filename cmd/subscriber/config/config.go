package config

import (
	"github.com/mmontes11/crypto-trade/internal/config"
	nats "github.com/nats-io/nats.go"
)

var (
	// Env indicates the environment name
	Env = config.GetEnv("ENV", "development")
	// NatsURL is the Nats instance used for publishing
	NatsURL = config.GetEnv("NATS_URL", nats.DefaultURL)
	// NumSubscribers is the number of trade subscribers
	NumSubscribers = config.GetIntEnv("NUM_SUBSCRIBERS", 4)
	// Subject is the subject for subscribing to trades
	Subject = config.GetEnv("SUBJECT", "trades")
	// ClickHouseURL it the URL of the ClickHouse instance
	ClickHouseURL = config.GetEnv("CLICKHOUSE_URL", "tcp://127.0.0.1:9000")
	// MigrationsURL is the URL of data migrations
	MigrationsURL = config.GetEnv("MIGRATIONS_URL", "file://model/migrations")
)

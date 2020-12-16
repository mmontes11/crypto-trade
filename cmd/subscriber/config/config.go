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
)

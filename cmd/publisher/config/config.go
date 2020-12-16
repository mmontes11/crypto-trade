package config

import (
	"time"

	"github.com/mmontes11/crypto-trade/internal/config"
	nats "github.com/nats-io/nats.go"
)

var (
	// Env indicates the environment name
	Env = config.GetEnv("ENV", "development")
	// NatsURL is the Nats instance used for publishing
	NatsURL = config.GetEnv("NATS_URL", nats.DefaultURL)
	// PublishInterval is the interval between publications
	PublishInterval = time.Duration(config.GetIntEnv("PUBLISH_INTERVAL_MS", 500)) * time.Millisecond
	// NumPublishers is the number of trade publishers
	NumPublishers = config.GetIntEnv("NUM_PUBLISHERS", 4)
	// Subject is the subject for publishing trades
	Subject = config.GetEnv("SUBJECT", "trades")
)

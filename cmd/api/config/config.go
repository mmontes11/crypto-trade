package config

import "github.com/mmontes11/crypto-trade/internal/config"

var (
	// Port indicates the API service port. It uses default K8s service port env variable
	Port = config.GetEnv("CRYPTO_TRADE_API_SERVICE_PORT", "10000")
	// Env indicates the environment name
	Env = config.GetEnv("ENV", "development")
	// ClickHouseURL it the URL of the ClickHouse instance
	ClickHouseURL = config.GetEnv("CLICKHOUSE_URL", "tcp://127.0.0.1:9000?username=crypto&password=crypto&database=crypto")
)

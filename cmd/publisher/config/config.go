package config

import "github.com/mmontes11/crypto-trade/internal/config"

var (
	// Env indicates the environment name
	Env = config.GetEnv("ENV", "development")
)

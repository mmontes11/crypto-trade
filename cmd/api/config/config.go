package config

import "github.com/mmontes11/crypto-trade/internal/config"

var (
	// Port indicates the API service port. It uses default K8s service port env variable
	Port = config.GetEnv("API_SERVICE_PORT", "10000")
	// Env indicates the environment name
	Env = config.GetEnv("ENV", "development")
)

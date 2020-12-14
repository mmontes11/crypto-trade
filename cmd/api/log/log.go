package log

import (
	"github.com/mmontes11/crypto-trade/cmd/api/config"
	"github.com/mmontes11/crypto-trade/internal/log"
)

// Logger instance
var Logger log.LoggerI

// Init initializes logger
func Init() {
	Logger = log.Init(config.Env)
}

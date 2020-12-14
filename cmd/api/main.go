package main

import (
	"github.com/mmontes11/crypto-trade/cmd/api/api"
	"github.com/mmontes11/crypto-trade/cmd/api/log"
)

func main() {
	log.Init()
	api.Init()
}

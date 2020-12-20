package main

import (
	"github.com/mmontes11/crypto-trade/cmd/api/api"
	"github.com/mmontes11/crypto-trade/cmd/api/config"
	"github.com/mmontes11/crypto-trade/cmd/api/controller"
	"github.com/mmontes11/crypto-trade/cmd/api/log"
	"github.com/mmontes11/crypto-trade/cmd/api/model"

	ch "github.com/mmontes11/crypto-trade/pkg/database/clickhouse"
)

func main() {
	log.Init()

	db, err := ch.Connect(config.ClickHouseURL)
	if err != nil {
		log.Logger.Fatal(err)
	}

	log.Logger.Infof("Connected to ClickHouse")

	tradeController := controller.NewTradeController(db, model.NewTradeRepository())

	api := api.NewAPI(tradeController)
	api.Listen()
}

package main

import (
	"github.com/mmontes11/crypto-trade/cmd/subscriber/config"
	"github.com/mmontes11/crypto-trade/cmd/subscriber/controller"
	"github.com/mmontes11/crypto-trade/cmd/subscriber/log"
	"github.com/mmontes11/crypto-trade/cmd/subscriber/model"
	ch "github.com/mmontes11/crypto-trade/pkg/clickhouse"
	nats "github.com/nats-io/nats.go"
)

func main() {
	log.Init()

	nc, err := nats.Connect(config.NatsURL)
	if err != nil {
		log.Logger.Fatal(err)
	}

	log.Logger.Info("Connected to NATS")

	db, err := ch.Connect(config.ClickHouseURL)
	if err != nil {
		log.Logger.Fatal(err)
	}

	log.Logger.Infof("Connected to ClickHouse")
	log.Logger.Info("Running migrations...")

	err = ch.MigrateUp(db, "file://model/migrations")
	if err != nil {
		log.Logger.Fatal(err)
	}

	log.Logger.Info("Migrations completed successfully")

	tradeController := controller.NewTradeController(nc, db, model.NewTradeRepository())
	tradeController.Subscribe()
}

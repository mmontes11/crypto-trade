package main

import (
	"github.com/mmontes11/crypto-trade/cmd/subscriber/config"
	"github.com/mmontes11/crypto-trade/cmd/subscriber/controller"
	"github.com/mmontes11/crypto-trade/cmd/subscriber/log"
	ch "github.com/mmontes11/crypto-trade/pkg/clickhouse"
	nats "github.com/nats-io/nats.go"
)

func main() {
	log.Init()

	nc, err := nats.Connect(config.NatsURL)
	if err != nil {
		log.Logger.Fatal(err)
	}

	log.Logger.Infof("Connected to NATS at %s", config.NatsURL)

	db, err := ch.Connect(config.ClickHouseURL)
	if err != nil {
		log.Logger.Fatal(err)
	}

	log.Logger.Infof("Connected to ClickHouse at %s", config.ClickHouseURL)
	log.Logger.Info("Running migrations...")

	err = ch.MigrateUp(db, "file://migrations")
	if err != nil {
		log.Logger.Fatal(err)
	}

	log.Logger.Info("Migrations completed successfully")

	subscribeController := controller.NewSubscribeController(nc, db)
	subscribeController.Subscribe()
}

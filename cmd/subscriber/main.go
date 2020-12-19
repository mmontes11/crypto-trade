package main

import (
	"github.com/mmontes11/crypto-trade/cmd/subscriber/config"
	"github.com/mmontes11/crypto-trade/cmd/subscriber/controller"
	"github.com/mmontes11/crypto-trade/cmd/subscriber/log"
	"github.com/mmontes11/crypto-trade/pkg/database"
	nats "github.com/nats-io/nats.go"
)

func main() {
	log.Init()

	nc, err := nats.Connect(config.NatsURL)
	if err != nil {
		log.Logger.Fatal(err)
	}

	log.Logger.Infof("Connected to NATS at %s", config.NatsURL)

	configDB := database.Config{
		URL:            config.ClickHouseURL,
		MigrationFiles: "file://migrations",
	}
	ch := database.NewClickHouse(configDB)

	err = ch.Connect()
	if err != nil {
		log.Logger.Fatal(err)
	}

	log.Logger.Infof("Connected to ClickHouse at %s", config.ClickHouseURL)
	log.Logger.Info("Running migrations...")

	err = ch.MigrateUp()
	if err != nil {
		log.Logger.Fatal(err)
	}

	log.Logger.Info("Migrations completed successfully")

	subscribeController := controller.NewSubscribeController(nc)
	subscribeController.Subscribe()
}

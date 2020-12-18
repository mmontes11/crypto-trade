package main

import (
	"database/sql"

	ch "github.com/ClickHouse/clickhouse-go"
	"github.com/golang-migrate/migrate/v4"
	migrateCH "github.com/golang-migrate/migrate/v4/database/clickhouse"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/mmontes11/crypto-trade/cmd/subscriber/config"
	"github.com/mmontes11/crypto-trade/cmd/subscriber/controller"
	"github.com/mmontes11/crypto-trade/cmd/subscriber/log"
	nats "github.com/nats-io/nats.go"
)

func main() {
	log.Init()

	nc, err := nats.Connect(config.NatsURL)
	if err != nil {
		log.Logger.Fatal(err)
	}

	log.Logger.Infof("Connected to NATS at %s", config.NatsURL)

	db, err := sql.Open("clickhouse", config.ClickHouseURL)
	if err != nil {
		log.Logger.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		if exception, ok := err.(*ch.Exception); ok {
			log.Logger.Errorf("[%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		} else {
			log.Logger.Error(err)
		}
		return
	}
	log.Logger.Infof("Connected to ClickHouse at %s", config.ClickHouseURL)

	log.Logger.Info("Running migrations...")
	driver, err := migrateCH.WithInstance(db, &migrateCH.Config{})
	if err != nil {
		log.Logger.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://migrations", "clickhouse", driver)
	if err != nil {
		log.Logger.Fatal(err)
	}

	m.Steps(1)
	log.Logger.Info("Migrations completed successfully")

	subscribeController := controller.NewSubscribeController(nc)
	subscribeController.Subscribe()
}

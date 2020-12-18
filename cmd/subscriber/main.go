package main

import (
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

	subscribeController := controller.NewSubscribeController(nc)
	subscribeController.Subscribe()
}

package main

import (
	"github.com/mmontes11/crypto-trade/cmd/publisher/config"
	"github.com/mmontes11/crypto-trade/cmd/publisher/log"
	"github.com/mmontes11/crypto-trade/internal/core"
	nats "github.com/nats-io/nats.go"
)

func main() {
	log.Init()

	nc, err := nats.Connect(config.NatsURL)
	if err != nil {
		log.Logger.Fatal(err)
	}
	c, _ := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	defer c.Close()

	subject := "trades"
	trade := core.NewRandTrade()

	log.Logger.Debugf("Publishing in \"%s\": \"%s\"", subject, trade)

	c.Publish(subject, trade)
}

package main

import (
	ctx "context"

	"github.com/mmontes11/crypto-trade/cmd/publisher/config"
	"github.com/mmontes11/crypto-trade/cmd/publisher/log"
	"github.com/mmontes11/crypto-trade/internal/core"
	"github.com/mmontes11/crypto-trade/pkg/scheduler"
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

	publish := func() {
		trade := core.NewRandTrade()
		log.Logger.Debugf("Publishing in \"%s\": \"%s\"", subject, trade)
		c.Publish(subject, core.NewRandTrade())
	}

	scheduler := scheduler.New(config.PublishInterval, publish)
	scheduler.Start(ctx.Background())
}

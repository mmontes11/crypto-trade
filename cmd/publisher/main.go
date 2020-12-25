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
	c, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		log.Logger.Fatal(err)
	}
	defer c.Close()

	log.Logger.Info("Connected to NATS")

	trades := make(chan core.Trade, config.NumPublishers)

	for i := 0; i < config.NumPublishers; i++ {
		go func(id int) {
			log.Logger.Infof("[Worker %d] Started publishing", id)
			for {
				select {
				case t := <-trades:
					log.Logger.Debugf("[Worker %d] Publishing in \"%s\": \"%s\"", id, config.Subject, t)
					c.Publish(config.Subject, t)
				}
			}
		}(i)
	}

	publish := func() {
		for i := 0; i < config.NumPublishers; i++ {
			trades <- core.NewRandTrade()
		}
	}

	scheduler := scheduler.New(config.PublishInterval, publish)
	scheduler.Start(ctx.Background())
}

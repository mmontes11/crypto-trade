package main

import (
	"github.com/mmontes11/crypto-trade/cmd/subscriber/config"
	"github.com/mmontes11/crypto-trade/cmd/subscriber/log"
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

	ch := make(chan *nats.Msg, 64)
	sub, err := nc.ChanSubscribe(subject, ch)
	if err != nil {
		log.Logger.Fatal(err)
	}
	defer func() {
		close(ch)
		sub.Unsubscribe()
	}()

	log.Logger.Debugf("Subscribed to \"%s\"...", subject)

	for msg := range ch {
		log.Logger.Debugf("Received: %s", msg)
	}
}

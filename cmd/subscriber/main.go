package main

import (
	"sync"

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
	queue := "workers"
	wg := sync.WaitGroup{}

	for i := 0; i < config.NumSubscribers; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			ch := make(chan *nats.Msg, 64)
			sub, err := nc.ChanQueueSubscribe(subject, queue, ch)
			if err != nil {
				log.Logger.Fatal(err)
			}
			defer func() {
				close(ch)
				sub.Unsubscribe()
			}()

			log.Logger.Debugf("[Worker %d] Subscribed to \"%s\"...", id, subject)

			for msg := range ch {
				log.Logger.Debugf("[Worker %d] Received: %s", id, msg)
			}
		}(i)
	}

	wg.Wait()
}

package controller

import (
	"database/sql"
	"sync"

	"github.com/mmontes11/crypto-trade/cmd/subscriber/config"
	"github.com/mmontes11/crypto-trade/cmd/subscriber/log"
	nats "github.com/nats-io/nats.go"
)

// SubscribeControllerI defines controller operations
type SubscribeControllerI interface {
	Subscribe()
}

// SubscribeController implements controller operations
type SubscribeController struct {
	natsConn *nats.Conn
	db       *sql.DB
}

// NewSubscribeController creates a new controller instance
func NewSubscribeController(natsConn *nats.Conn, db *sql.DB) SubscribeControllerI {
	return &SubscribeController{
		natsConn,
		db,
	}
}

// Subscribe subscribe to trading subject
func (sc *SubscribeController) Subscribe() {
	c, _ := nats.NewEncodedConn(sc.natsConn, nats.JSON_ENCODER)
	defer c.Close()

	queue := "workers"
	wg := sync.WaitGroup{}

	for i := 0; i < config.NumSubscribers; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			ch := make(chan *nats.Msg, 64)
			sub, err := sc.natsConn.ChanQueueSubscribe(config.Subject, queue, ch)
			if err != nil {
				log.Logger.Fatal(err)
			}
			defer func() {
				close(ch)
				sub.Unsubscribe()
			}()

			log.Logger.Debugf("[Worker %d] Subscribed to \"%s\"...", id, config.Subject)

			for msg := range ch {
				log.Logger.Debugf("[Worker %d] Received: %s", id, msg)
			}
		}(i)
	}

	wg.Wait()
}

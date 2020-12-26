package controller

import (
	ctx "context"
	"database/sql"
	"encoding/json"
	"sync"

	"github.com/mmontes11/crypto-trade/cmd/subscriber/config"
	"github.com/mmontes11/crypto-trade/cmd/subscriber/log"
	"github.com/mmontes11/crypto-trade/cmd/subscriber/model"
	"github.com/mmontes11/crypto-trade/internal/core"
	"github.com/mmontes11/crypto-trade/pkg/database"
	nats "github.com/nats-io/nats.go"
)

// TradeControllerI defines controller operations
type TradeControllerI interface {
	Subscribe()
}

// TradeController implements controller operations
type TradeController struct {
	natsConn   *nats.Conn
	db         *sql.DB
	repository model.TradeRepositoryI
}

// NewTradeController creates a new controller instance
func NewTradeController(natsConn *nats.Conn, db *sql.DB, r model.TradeRepositoryI) TradeControllerI {
	return &TradeController{
		natsConn,
		db,
		r,
	}
}

// Subscribe subscribe to trading subject
func (tc *TradeController) Subscribe() {
	c, _ := nats.NewEncodedConn(tc.natsConn, nats.JSON_ENCODER)
	defer c.Close()

	queue := "workers"
	wg := sync.WaitGroup{}

	for i := 0; i < config.NumSubscribers; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			ch := make(chan *nats.Msg, 64)
			sub, err := tc.natsConn.ChanQueueSubscribe(config.Subject, queue, ch)
			if err != nil {
				log.Logger.Fatal(err)
			}
			defer func() {
				close(ch)
				sub.Unsubscribe()
			}()

			log.Logger.Infof("[Worker %d] Subscribed to \"%s\"...", id, config.Subject)

			for msg := range ch {
				log.Logger.Debugf("[Worker %d] Received: %s", id, msg.Data)
				tc.handleMessage(msg)
			}
		}(i)
	}

	wg.Wait()
}

func (tc *TradeController) handleMessage(msg *nats.Msg) {
	var t core.Trade
	err := json.Unmarshal(msg.Data, &t)

	if err != nil {
		log.Logger.Error(err)
		return
	}

	err = tc.saveTrade(t)
	if err != nil {
		log.Logger.Error(err)
	}
}

func (tc *TradeController) saveTrade(t core.Trade) error {
	txFn := func(ctx ctx.Context, tx *sql.Tx) error {
		return tc.repository.SaveTrade(ctx, tx, t)
	}

	return database.Tx(ctx.Background(), tc.db, txFn)
}

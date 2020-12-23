package core

import (
	"math/rand"
	"time"
)

// Trade cryptocurrency trade
type Trade struct {
	Time  time.Time      `json:"time"`
	Side  string         `json:"side"`
	Size  CurrencyAmount `json:"size"`
	Price CurrencyAmount `json:"price"`
}

// CurrencyAmount quantifies the amount of a specific currency
type CurrencyAmount struct {
	Amount   float32 `json:"amount"`
	Currency string  `json:"currency"`
}

// TradeParams used in trade related methods
type TradeParams struct {
	GroupBy  string
	Crypto   string
	Currency string
	Limit    int
}

// Validate checks if params are valid
func (tp *TradeParams) Validate() error {
	if tp.GroupBy != "" {
		supportedGroupBy := false
		for _, v := range []string{"second", "minute", "hour"} {
			if v == tp.GroupBy {
				supportedGroupBy = true
			}
		}
		if !supportedGroupBy {
			return &ErrInvalidField{"groupBy"}
		}
	}
	if tp.Crypto == "" {
		return &ErrInvalidField{"crypto"}
	}
	if tp.Currency == "" {
		return &ErrInvalidField{"currency"}
	}
	if tp.Limit == 0 {
		return &ErrInvalidField{"limit"}
	}
	return nil
}

const (
	btc = "btc"
	eth = "eth"
	usd = "usd"
	eur = "eur"
)

// NewRandTrade creates a random Trade
func NewRandTrade() Trade {
	size := randSize()
	price := randPrice(size.Currency)

	return Trade{
		Time:  time.Now(),
		Side:  randSide(),
		Size:  size,
		Price: price,
	}
}

func randSide() string {
	if randBool() {
		return "sell"
	}
	return "buy"
}

func randSize() CurrencyAmount {
	size := randFloat(0.1, 10)
	var currency string
	if randBool() {
		currency = btc
	} else {
		currency = eth
	}

	return CurrencyAmount{
		Amount:   size,
		Currency: currency,
	}
}

func randPrice(crypto string) CurrencyAmount {
	var amount float32
	switch crypto {
	case btc:
		amount = randFloat(18000, 22000)
	case eth:
		amount = randFloat(500, 800)
	default:
		panic("Unsupported cryptocurrency: " + crypto)
	}

	var currency string
	if randBool() {
		currency = usd
	} else {
		currency = eur
	}

	return CurrencyAmount{
		Amount:   amount,
		Currency: currency,
	}
}

func randFloat(min, max float32) float32 {
	return min + rand.Float32()*(max-min)
}

func randBool() bool {
	return rand.Intn(2) == 0
}

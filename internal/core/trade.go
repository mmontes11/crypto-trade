package core

import (
	"math/rand"
	"time"
)

// Trade cryptocurrency trade
type Trade struct {
	Time       time.Time  `json:"time"`
	Side       string     `json:"side"`
	CryptoSize CryptoSize `json:"cryptoSize"`
	Price      Price      `json:"price"`
}

// CryptoSize is the number of cryptocurrencies traded
type CryptoSize struct {
	Size     float64 `json:"size"`
	Currency string  `json:"currency"`
}

// Price is the price of an unit of the cryptocurrency traded
type Price struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}

// TradeParams used in trade related methods
type TradeParams struct {
	Crypto   string
	Currency string
	Side     string
	Limit    int
}

// Validate checks if params are valid
func (tp *TradeParams) Validate() error {
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
		Time:       time.Now(),
		Side:       randSide(),
		CryptoSize: size,
		Price:      price,
	}
}

func randSide() string {
	if randBool() {
		return "sell"
	}
	return "buy"
}

func randSize() CryptoSize {
	size := randFloat(0.1, 10)
	var currency string
	if randBool() {
		currency = btc
	} else {
		currency = eth
	}

	return CryptoSize{
		Size:     size,
		Currency: currency,
	}
}

func randPrice(crypto string) Price {
	var amount float64
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

	return Price{
		Amount:   amount,
		Currency: currency,
	}
}

func randFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func randBool() bool {
	return rand.Intn(2) == 0
}

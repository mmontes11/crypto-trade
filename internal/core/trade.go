package core

import (
	"math/rand"
	"time"
)

// Trade cryptocurrency trade
type Trade struct {
	Time time.Time `json:"time"`
	// Buy or sell
	Side string `json:"side"`
	// Number of cryptocurrencies traded
	Size TradePrice `json:"size"`
	// Price of the traded cryptocurrencies
	Price TradePrice `json:"price"`
}

// TradePrice price of a (crypto)currency expressed in other currency
type TradePrice struct {
	Price    float64 `json:"price"`
	Currency string  `json:"currency"`
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
	price := randPrice(size)

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

func randSize() TradePrice {
	price := randFloat(0.1, 10)
	var currency string
	if randBool() {
		currency = btc
	} else {
		currency = eth
	}

	return TradePrice{
		Price:    price,
		Currency: currency,
	}
}

func randPrice(size TradePrice) TradePrice {
	cryptoUnitPrice := randCryptoUnitPrice(size.Currency)
	price := size.Price * cryptoUnitPrice.Price

	return TradePrice{
		Price:    price,
		Currency: cryptoUnitPrice.Currency,
	}
}

func randCryptoUnitPrice(crypto string) TradePrice {
	var price float64
	switch crypto {
	case btc:
		price = randFloat(18000, 22000)
	case eth:
		price = randFloat(500, 800)
	}

	var currency string
	if randBool() {
		currency = usd
	} else {
		currency = eur
	}

	return TradePrice{
		Price:    price,
		Currency: currency,
	}
}

func randFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func randBool() bool {
	return rand.Intn(2) == 0
}

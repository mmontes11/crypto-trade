package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/mmontes11/crypto-trade/cmd/api/controller"
	"github.com/mmontes11/crypto-trade/cmd/api/log"
	"github.com/mmontes11/crypto-trade/internal/core"
)

// Handler implements HTTP handlerw
type Handler struct {
	controller controller.TradeControllerI
}

var (
	errTradesNotFound = errors.New("No trades found")
)

// NewHandler creates a new handler
func NewHandler(ctrl controller.TradeControllerI) Handler {
	return Handler{
		ctrl,
	}
}

func (h Handler) healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (h Handler) tradesHandler(w http.ResponseWriter, r *http.Request) {
	params, err := getTradeParams(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	trades, err := h.controller.GetTrades(params)
	if err != nil {
		log.Logger.Error(err)
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
	if len(trades) == 0 {
		http.Error(w, errTradesNotFound.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(trades)
}

func getTradeParams(r *http.Request) (core.TradeParams, error) {
	q := r.URL.Query()
	groupBy := q.Get("groupBy")
	crypto := q.Get("crypto")
	currency := q.Get("currency")
	limit, err := parseIntParam(q.Get("limit"), 100)

	if err != nil {
		return core.TradeParams{}, err
	}

	params := core.TradeParams{
		GroupBy:  groupBy,
		Crypto:   crypto,
		Currency: currency,
		Limit:    limit,
	}

	return params, params.Validate()
}

func parseIntParam(paramStr string, defaultValue int) (int, error) {
	if len(paramStr) == 0 {
		return defaultValue, nil
	}
	param, err := strconv.Atoi(paramStr)
	if err != nil {
		return 0, err
	}
	if param < 0 {
		return 0, errors.New("Param cannot not be negative")
	}
	return param, nil
}

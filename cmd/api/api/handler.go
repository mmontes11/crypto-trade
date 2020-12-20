package api

import (
	"net/http"

	"github.com/mmontes11/crypto-trade/cmd/api/controller"
)

// Handler implements HTTP handlerw
type Handler struct {
	controller controller.TradeControllerI
}

// NewHandler creates a new handler
func NewHandler(ctrl controller.TradeControllerI) Handler {
	return Handler{
		ctrl,
	}
}

func (h Handler) healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

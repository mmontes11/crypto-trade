package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mmontes11/crypto-trade/cmd/api/config"
	"github.com/mmontes11/crypto-trade/cmd/api/controller"
	"github.com/mmontes11/crypto-trade/cmd/api/log"
)

// API implements a REST API
type API struct {
	handler Handler
}

// NewAPI creates a new API
func NewAPI(ctrl controller.TradeControllerI) API {
	handler := NewHandler(ctrl)
	return API{
		handler,
	}
}

// Listen starts a new HTTP server
func (a API) Listen() {
	router := a.createRouter()
	addr := fmt.Sprintf(":%s", config.Port)

	log.Logger.Info("Server listening on port ", config.Port)
	log.Logger.Info("Environment: ", config.Env)

	err := http.ListenAndServe(addr, router)
	if err != nil {
		log.Logger.Fatal(err)
	}
}

func (a API) createRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/health", a.handler.healthHandler)

	api := router.PathPrefix("/api").Subrouter()
	api.HandleFunc("/trades", a.handler.tradesHandler).Methods(http.MethodGet)

	return router
}

package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mmontes11/crypto-trade/cmd/api/config"
	"github.com/mmontes11/crypto-trade/cmd/api/log"
)

// Init starts a new HTTP server
func Init() {
	router := createRouter()
	addr := fmt.Sprintf(":%s", config.Port)

	log.Logger.Info("Server listening on port ", config.Port)
	log.Logger.Info("Environment: ", config.Env)

	err := http.ListenAndServe(addr, router)
	if err != nil {
		log.Logger.Fatal(err)
	}
}

func createRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/health", healthHandler)

	return router
}

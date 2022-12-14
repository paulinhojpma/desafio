package web

import (
	mux "github.com/gorilla/mux"
	// newrelic "github.com/newrelic/go-agent"

	"backend/desafio/core"
)

var (
	ExcludeRoutes []string
)

// Router ...
func Router(h *Handler) *mux.Router {
	ExcludeRoutes = make([]string, 0)
	router := core.Router()
	router.HandleFunc("/transactions", h.CreateTransactions).Methods("POST")
	router.HandleFunc("/transactions", h.ListarTransacoes).Methods("GET")

	excludeRoutes()

	return router
}

func excludeRoutes() {

}

// RouterTest ...
func RouterTest(h *Handler) *mux.Router {
	router := core.Router()
	router.HandleFunc("/transactions", h.CreateTransactions).Methods("POST")
	router.HandleFunc("/transactions", h.ListarTransacoes).Methods("GET")

	return router
}

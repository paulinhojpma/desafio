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
	// router.HandleFunc("/transacoes", h.CadastrarPlaneta).Methods("POST")
	// router.HandleFunc("/transacoes", h.ListarPlanetas).Methods("GET")

	excludeRoutes()

	return router
}

func excludeRoutes() {

}

// RouterTest ...
func RouterTest(h *Handler) *mux.Router {
	router := core.Router()
	// router.HandleFunc("/planetas", h.CadastrarPlaneta).Methods("POST")
	// router.HandleFunc("/planetas", h.ListarPlanetas).Methods("GET")
	// router.HandleFunc("/planetas/{idPlaneta}", h.GetPlaneta).Methods("GET")
	// router.HandleFunc("/planetas/{idPlaneta}", h.DeletePlaneta).Methods("DELETE")

	return router
}

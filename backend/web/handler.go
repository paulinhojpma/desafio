package web

import "backend/desafio/database"

// import (
// 	// "errors"
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"strconv"
// 	"time"

// 	"backend/desafio/core"
// 	"backend/desafio/database"

// 	mux "github.com/gorilla/mux"
// 	// "sw-sys/api-service/logger"
// )

// var (
// 	timeDefaultCache = time.Minute * 10
// )

// const (
// 	URLSWAPI = "https://swapi.dev/api/planets/?search="
// )

// // Handler ...
type Handler struct {
	// Relic        newrelic.Application
	// ClientRedis  *redis.Client
	// DB            *database.DataBase

	Database *database.IDataBase
	// Message       *messaging.IMessageClient
	// Cache         *cache.ICacheClient
	// Logger        *logger.ILogger
	// Storage       *storage.IStorage
	// Auth          *auth.IAuth
	// Broker        *Broker
	// Authorization *Authorization
	// Email         *email.Email
}

// type ResponseApiStarWars struct {
// 	Count     int              `json:"count"`
// 	PlanetsSW []PlanetStarWars `json:"results"`
// }
// type PlanetStarWars struct {
// 	Films []string `json:"films"`
// }

// func (h *Handler) HandleOption(w http.ResponseWriter, r *http.Request) {
// 	// w.WriteHeader(http.StatusOK)
// 	log.Println("ENTROU NO OPTIONS")
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.Header().Set("Access-Control-Allow-Headers", "*")
// 	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

// 	// (*w).Header().Set("Access-Control-Allow-Origin", "*")
// 	//   (*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
// 	//   (*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

// 	return
// }

// // CadastrarPlaneta handler que recebe a requisição e executa o cadastro do novo planeta
// func (h *Handler) CadastrarPlaneta(w http.ResponseWriter, r *http.Request) {
// 	planet := &Planet{}
// 	decoder := json.NewDecoder(r.Body)
// 	err := decoder.Decode(&planet)
// 	if err != nil {
// 		h.CoreRespondErro(w, r, "", ErrorFormatoRequest, ErrCadPlaneta, http.StatusBadRequest)
// 		return
// 	}
// 	resp, errGet := http.Get(fmt.Sprintf("%s%s", URLSWAPI, planet.Nome))
// 	if errGet != nil {
// 		h.CoreRespondErro(w, r, "", ErrorInternalServer, ErrCadPlaneta, http.StatusBadRequest)
// 		return
// 	}
// 	respSWAPI := &ResponseApiStarWars{}
// 	decoder = json.NewDecoder(resp.Body)
// 	err = decoder.Decode(&respSWAPI)
// 	if err != nil {
// 		h.CoreRespondErro(w, r, "", ErrorInternalServer, ErrCadPlaneta, http.StatusBadRequest)
// 		return
// 	}
// 	if respSWAPI.Count == 0 {
// 		h.CoreRespondErro(w, r, "", ErrorPlanetaNEncontrando, ErrCadPlaneta, http.StatusBadRequest)
// 		return
// 	}
// 	aparicoes := len(respSWAPI.PlanetsSW[0].Films)
// 	log.Printf("Aparições do filme %s: %d", planet.Nome, aparicoes)
// 	planet.Aparicoes = aparicoes
// 	errCad := Cadastrar(h, planet)
// 	if errCad != nil {
// 		log.Println("ERRO ao cadastrar: ", errCad)
// 		h.CoreRespondErro(w, r, "", ErrorInternalServer, ErrCadPlaneta, http.StatusBadRequest)
// 		return
// 	}

// 	response := &PlanetResponse{}
// 	response.CodResposta = http.StatusCreated
// 	response.Mensagem = CadPlaneta
// 	response.Planet = planet

// 	h.CoreRespondSucess(w, r, http.StatusCreated, response)
// 	return

// }

// // GetPlaneta handler que recebe a requisição e retorna o Planeta de acordo com seu ID
// func (h *Handler) GetPlaneta(w http.ResponseWriter, r *http.Request) {

// 	vars := mux.Vars(r)
// 	ID := vars["idPlaneta"]

// 	planet, err := Get(h, ID)
// 	if err != nil {
// 		log.Println("ERRO ao pegar: ", err)
// 		h.CoreRespondErro(w, r, "", ErrorPlanetaNEncontrando, ErrGetPlaneta, http.StatusBadRequest)
// 		return
// 	}

// 	response := &PlanetResponse{}
// 	response.CodResposta = http.StatusOK
// 	response.Mensagem = GetPlaneta
// 	response.Planet = planet

// 	h.CoreRespondSucess(w, r, http.StatusOK, response)
// 	return

// }

// // ListarPlaneta handler que recebe a requisição e lista todos os planetas, caso venha com o
// // parâmetro Query faz a consulta pelo nome o planeta.
// func (h *Handler) ListarPlanetas(w http.ResponseWriter, r *http.Request) {
// 	planets := []Planet{}
// 	// vars := mux.Vars(r)
// 	// ID := vars["idPlaneta"]
// 	nome := ""
// 	if len(r.URL.Query()["nome"]) > 0 {
// 		nome = r.URL.Query()["nome"][0]
// 	}
// 	log.Println("Query retornada: ", nome)
// 	if nome == "" {
// 		log.Println()
// 		pls, err := Listar(h)
// 		if err != nil {

// 			h.CoreRespondErro(w, r, "", ErrorPlanetaNEncontrando, ErrListarPlaneta, http.StatusBadRequest)
// 			return
// 		}
// 		planets = pls
// 	} else {
// 		pls, err := Consultar(h, nome)
// 		if err != nil {

// 			h.CoreRespondErro(w, r, "", ErrorPlanetaNEncontrando, ErrListarPlaneta, http.StatusBadRequest)
// 			return
// 		}
// 		planets = pls
// 	}

// 	response := &PlanetsResponse{}
// 	response.CodResposta = http.StatusOK
// 	response.Mensagem = ListPlanetas
// 	response.Planets = planets

// 	h.CoreRespondSucess(w, r, http.StatusOK, response)
// 	return

// }

// // DeltePlaneta handler que recebe a requisição e remove de acordo com seu ID
// func (h *Handler) DeletePlaneta(w http.ResponseWriter, r *http.Request) {

// 	vars := mux.Vars(r)
// 	ID := vars["idPlaneta"]

// 	err := Delete(h, ID)
// 	if err != nil {
// 		log.Println("ERRO ao pegar: ", err)
// 		h.CoreRespondErro(w, r, "", ErrorPlanetaNEncontrando, ErrDeletePlaneta, http.StatusBadRequest)
// 		return
// 	}

// 	response := &PlanetResponse{}
// 	response.CodResposta = http.StatusOK
// 	response.Mensagem = DeletePlaneta

// 	h.CoreRespondSucess(w, r, http.StatusOK, response)
// 	return

// }

// //CoreRespondErro ...
// /**
// Responsavel por centralizar as chamadas de core.RespondErro ...
// */
// func (h *Handler) CoreRespondErro(w http.ResponseWriter, r *http.Request, idOperation string, erro error, message string, codeError int) {
// 	core.Respond(w, r, codeError, core.ErrDetail{
// 		Resource: message, Code: strconv.Itoa(codeError), Message: erro.Error(), IDOperation: idOperation,
// 	})
// }

// func (h *Handler) CoreRespondSucess(w http.ResponseWriter, r *http.Request, code int, object interface{}) {
// 	core.Respond(w, r, code, object)
// }

package main

import (
	"backend/desafio/config"
	"backend/desafio/database"
	"backend/desafio/web"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/rs/cors"
)

const (
	TimeOutSecond = 120
)

func main() {

	configNew := config.Config()
	stage := os.Getenv("STAGE")
	log.Println("ENVIROMENT - ", stage)
	if stage == "" {
		configNew = config.NewConfig("conf.json")
	} else {
		configNew = config.NewConfig("")
	}
	handler := &web.Handler{}

	DBOpt := &database.OptionsDBClient{
		URL:    configNew.DBHost,
		Driver: configNew.DBDriver,
	}

	clientDB, errDB := DBOpt.ConfiguraDatabase()
	if errDB != nil {
		log.Println(errDB)
		os.Exit(1)
	}
	handler.Database = clientDB
	allowedParam := make(map[string][]string)
	if err := json.Unmarshal([]byte(`{"Origins":["*"],"Headers":["*"],"Methods":["GET","POST","PUT", "DELETE","OPTIONS"]}`), &allowedParam); err != nil {
		log.Println("Erro no json Unmarshal do allowedOrigins. Detalhe:", err)
		os.Exit(1)
	}

	router := web.Router(handler)
	c := cors.New(cors.Options{
		AllowedOrigins: allowedParam["Origins"],
		AllowedHeaders: allowedParam["Headers"],
		AllowedMethods: allowedParam["Methods"],
		// OptionsPassthrough: true,
		Debug: false,
	})
	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", configNew.Port),
		Handler:      c.Handler(router),
		ReadTimeout:  TimeOutSecond * time.Second,
		WriteTimeout: TimeOutSecond * time.Second,
	}
	log.Println("Esperando conexão")
	if err := s.ListenAndServe(); err != nil {
		log.Println("Erro ao iniciar Server. Erro:", err)
		os.Exit(1)
	}
}

package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	env "github.com/caarlos0/env"
)

var config *Configuracoes

//Configuracoes ...
type Configuracoes struct {
	Service string `json:"service" env:"SERVICE"`

	Port int `json:"port" env:"PORT"`

	AllowedParam string `json:"allowedParam" env:"ALLOWED_PARAM"`

	DBNome   string `json:"db-nome" env:"DB_NOME"`
	DBDriver string `json:"db-driver" env:"DB_DRIVER"`
	DBHost   string `json:"db-host" env:"DB_HOST"`
}

//NewConfig ...
func NewConfig(file string) *Configuracoes {
	var erro error

	conf := &Configuracoes{}

	if file != "" {
		fmt.Println(file)

		bufConf, err := ioutil.ReadFile(file)
		if err == nil {
			erro = json.Unmarshal(bufConf, conf)
			if erro != nil {
				log.Println("erro ao configurar -", erro)
			}
		}
	}

	//variaveis de ambiente sobrescrevem informacoes do json
	if erro = env.Parse(conf); erro != nil {
		log.Println(erro)
	}

	return conf
}

//Config ...
func Config() *Configuracoes {
	return config
}

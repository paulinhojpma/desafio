package database

import (
	"log"
	// "time"
)

// Planet monta a estrututra do plano

// IDataBase ..
type IDataBase interface {
	connectService(config *OptionsDBClient) error
	GetTransacao(key interface{}) (Produtor, error)
	CreateTransacao(produtor *Produtor) error
	ListTransacao() ([]Produtor, error)
}

// OptionsCacheClient ..
type OptionsDBClient struct {
	URL    string `json:"url"`
	DBName string `json:"DBName"`
	Driver string
}

// ConfiguraCache
func (o *OptionsDBClient) ConfiguraDatabase() (*IDataBase, error) {
	log.Println("Entrou no configura dataBase")
	var client IDataBase
	switch o.Driver {
	case "postgres":
		gormDB := &gormPostgres{}
		errGorm := gormDB.connectService(o)
		if errGorm != nil {
			return nil, errGorm
		}
		client = gormDB
	}
	return &client, nil
}

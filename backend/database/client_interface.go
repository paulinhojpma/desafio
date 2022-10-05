package database

import (
	"log"
	"time"
	// "time"
)

// Planet monta a estrututra do plano

type Transacao struct {
	ID      int
	Tipo    string
	Data    time.Time
	Produto string
	Valor   float64
}

func (Transacao) TableName() string {
	return "TRANSACAO"
}

// IDataBase ..
type IDataBase interface {
	connectService(config *OptionsDBClient) error
	GetPlanet(key interface{}) (Planet, error)
	CreatePlanet(value interface{}) error
	ListPlanets() ([]Planet, error)
	DeletePlanet(key interface{}) error
	UpdatePlanet(planet Planet) error
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

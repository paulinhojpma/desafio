package database

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func initDataBase() (IDataBase, error) {
	DBOpt := &OptionsDBClient{
		URL:    "postgresql://root:luke@localhost:5434/app",
		Driver: "postgres",
	}
	db, err := DBOpt.ConfiguraDatabase()
	return *db, err

}

func TestInitConnection(t *testing.T) {
	_, err := initDataBase()

	if err != nil {
		t.Error("Expect nil, got - ", err)
	}
}

func TestCreateTransacao(t *testing.T) {
	db, _ := initDataBase()

	trans := &Transacao{
		Tipo: TipoTransacao{
			Tipo: 4,
		},
		TipoID:  1,
		Data:    time.Now(),
		Produto: "outro",
		Valor:   6.8,
	}

	prod := &Produtor{
		Nome: "klebernilton",
	}

	prod.Transacoes = []Transacao{*trans}
	err := db.CreateTransacao([]*Produtor{prod})
	if err != nil {
		t.Error("Expect nil, got ", err)
	}

}

func TestGetTransacao(t *testing.T) {
	db, _ := initDataBase()

	trans, err := db.GetTransacao(4)
	byt, err := json.Marshal(trans)
	fmt.Println(string(byt))
	if err == nil {
		t.Error("Expect nil, got ", err)
	}

}

func TestListTransacao(t *testing.T) {
	db, _ := initDataBase()

	trans, err := db.ListTransacao()
	byt, err := json.Marshal(trans)
	fmt.Println(string(byt))
	if err == nil {
		t.Error("Expect nil, got ", err)
	}

}

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
	db, err := DBOpt.ConfigDatabase()
	return *db, err

}

func TestInitConnection(t *testing.T) {
	_, err := initDataBase()

	if err != nil {
		t.Error("Expect nil, got - ", err)
	}
}

func TestCreateTransaction(t *testing.T) {
	db, _ := initDataBase()

	trans := &Transaction{
		Type: TypeTransaction{
			Type: 4,
		},
		TypeID:  1,
		Date:    time.Now(),
		Product: "outro",
		Value:   6.8,
	}

	prod := &Producer{
		Name: "klebernilton",
	}

	prod.Transactions = []Transaction{*trans}
	err := db.CreateTransaction([]*Producer{prod})
	if err != nil {
		t.Error("Expect nil, got ", err)
	}

}

func TestGetTransaction(t *testing.T) {
	db, _ := initDataBase()

	trans, err := db.GetTransaction(4)
	byt, err := json.Marshal(trans)
	fmt.Println(string(byt))
	if err != nil {
		t.Error("Expect nil, got ", err)
	}

}

func TestListTransaction(t *testing.T) {
	db, _ := initDataBase()

	trans, err := db.ListTransaction()
	byt, err := json.Marshal(trans)
	fmt.Println(string(byt))
	if err != nil {
		t.Error("Expect nil, got ", err)
	}

}

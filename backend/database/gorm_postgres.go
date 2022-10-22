package database

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type gormPostgres struct {
	db *gorm.DB
}

type Transaction struct {
	gorm.Model
	//ID         int
	TypeID     int
	Type       TypeTransaction `gorm:"foreignKey:TypeID"`
	Date       time.Time
	Product    string
	Value      float64
	ProducerID int
}

type TypeTransaction struct {
	Type        int `gorm:"primaryKey"`
	Description string
	Nature      string
}

type Producer struct {
	ID           int
	Name         string
	Transactions []Transaction
	SumTotal     float64 `gorm:"-"`
}

func (Transaction) TableName() string {
	return "TRANSACTION"
}

func (TypeTransaction) TableName() string {
	return "TYPE_TRANSACTION"
}

func (Producer) TableName() string {
	return "PRODUCER"
}
func (gm *gormPostgres) connectService(config *OptionsDBClient) error {
	i := 0
	fmt.Println("DB URL - ", config.URL)
	for {
		db, err := gorm.Open(postgres.Open(config.URL), &gorm.Config{})
		if err != nil {
			if i >= 4 {
				return err
			} else {
				i++
			}
			time.Sleep(time.Second * 5)

		} else {
			gm.db = db
			log.Println("Connect on DataBase")
			break
		}

	}

	return nil
}

func (gm *gormPostgres) GetTransaction(key interface{}) (Producer, error) {
	produtor := &Producer{}
	result := gm.db.Preload("Transactions").Preload("Transactions.Type").First(produtor, key.(int))
	return *produtor, result.Error
}

func (gm *gormPostgres) CreateTransaction(produtores []*Producer) error {
	tx := gm.db.Begin()
	for _, produtor := range produtores {
		if err := gm.createProdutor(produtor, tx); err != nil {
			tx.Rollback()
			return err
		}
	}

	defer tx.Commit()
	return nil
}

func (gm *gormPostgres) ListTransaction() ([]Producer, error) {
	produtores := make([]Producer, 0)
	result := gm.db.Preload("Transactions").Preload("Transactions.Type").Find(&produtores)
	return produtores, result.Error

}

func (gm *gormPostgres) createProdutor(produtor *Producer, tx *gorm.DB) error {
	prd := &Producer{}
	err := gm.db.Where("name = ?", produtor.Name).First(prd).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return err
		}

	}
	if prd.ID > 0 {
		produtor.ID = prd.ID
		return nil
	}

	return tx.Create(produtor).Error

}

// func (gm *gormPostgres) getProdutor(tipoID int)(TypeTransaction, error){

// 		tipoTrans := &TypeTransaction{}

// 		gm.db.G

// }

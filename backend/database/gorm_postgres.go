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

type Transacao struct {
	gorm.Model
	//ID         int
	TipoID     int
	Tipo       TipoTransacao `gorm:"foreignKey:TipoID"`
	Data       time.Time
	Produto    string
	Valor      float64
	ProdutorID int
}

type TipoTransacao struct {
	Tipo      int `gorm:"primaryKey"`
	Descricao string
	Natureza  string
}

type Produtor struct {
	ID         int
	Nome       string
	Transacoes []Transacao
}

func (Transacao) TableName() string {
	return "TRANSCAO"
}

func (TipoTransacao) TableName() string {
	return "TIPO_TRANSACAO"
}

func (Produtor) TableName() string {
	return "PRODUTOR"
}
func (gm *gormPostgres) connectService(config *OptionsDBClient) error {
	i := 0

	for {
		db, err := gorm.Open(postgres.Open(config.URL), &gorm.Config{})
		if err != nil {
			if i >= 3 {
				return err
			} else {
				i++
			}
			time.Sleep(1000 * time.Millisecond)

		} else {
			gm.db = db
			log.Println("Conectou ao DB")
			break
		}

	}

	return nil
}

func (gm *gormPostgres) GetTransacao(key interface{}) (Produtor, error) {
	produtor := &Produtor{}
	result := gm.db.Preload("Transacoes").Preload("Transacoes.Tipo").First(produtor, key.(int))
	return *produtor, result.Error
}
func (gm *gormPostgres) CreateTransacao(produtor *Produtor) error {

	if err := gm.createProdutor(produtor); err != nil {
		return err
	}

	return nil
}

func (gm *gormPostgres) ListTransacao() ([]Produtor, error) {
	produtores := make([]Produtor, 0)
	result := gm.db.Preload("Transacoes").Preload("Transacoes.Tipo").Find(&produtores)
	return produtores, result.Error

}

func (gm *gormPostgres) createProdutor(produtor *Produtor) error {
	prd := &Produtor{}
	err := gm.db.Where("nome = ?", produtor.Nome).First(prd).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return err
		}

	}
	if prd.ID > 0 {
		produtor.ID = prd.ID
		fmt.Println("ID PRD ", produtor.ID)
		return nil
	}

	return gm.db.Create(produtor).Error

}

// func (gm *gormPostgres) getProdutor(tipoID int)(TipoTransacao, error){

// 		tipoTrans := &TipoTransacao{}

// 		gm.db.G

// }

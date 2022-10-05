package database

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type gormPostgres struct {
	db *gorm.DB
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

func (gm *gormPostgres) GetPlanet(key interface{}) (Planet, error) {
	planet := &Planet{}
	result := gm.db.First(planet, key.(int))

	return *planet, result.Error
}
func (gm *gormPostgres) CreatePlanet(value interface{}) error {
	result := gm.db.Create(value)
	return result.Error
}

func (gm *gormPostgres) ListPlanets() ([]Planet, error) {
	planets := make([]Planet, 0)
	result := gm.db.Find(&planets)
	return planets, result.Error

}

func (gm *gormPostgres) UpdatePlanet(planet Planet) error {

	result := gm.db.Model(&planet).Updates(planet)
	return result.Error
}
func (gm *gormPostgres) DeletePlanet(key interface{}) error {
	result := gm.db.Delete(&Planet{}, key)
	return result.Error
}

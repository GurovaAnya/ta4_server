package repository

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"log"
	"ta4/mod/model"
)

type TriggerRepository struct{}

func (repo TriggerRepository) CreateTrigger(Trigger model.TriggerModel) model.TriggerModel {
	db, err := gorm.Open("postgres", "user=postgres dbname=ta4 sslmode=disable password=docker")
	if err != nil {
		log.Panic(err)
	}
	db.Create(&Trigger)
	return Trigger
}

func (repo TriggerRepository) CreateTriggerToEventRelation(Trigger model.TriggerToEventModel) model.TriggerToEventModel {
	db, err := gorm.Open("postgres", "user=postgres dbname=ta4 sslmode=disable password=docker")
	if err != nil {
		log.Panic(err)
	}
	db.Create(&Trigger)
	return Trigger
}

func (repo TriggerRepository) GetTrigger(id uuid.UUID) model.TriggerModel {
	db, err := gorm.Open("postgres", "user=postgres dbname=ta4 sslmode=disable password=docker")
	if err != nil {
		log.Panic(err)
	}

	var result = model.TriggerModel{Id: id}
	db.Find(&result, result)
	return result
}

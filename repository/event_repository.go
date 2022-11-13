package repository

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"log"
	"ta4/mod/model"
)

type EventRepository struct{}

func (repo EventRepository) CreateEvent(event model.EventModel) model.EventModel {
	db, err := gorm.Open("postgres", "user=postgres dbname=ta4 sslmode=disable password=docker")
	if err != nil {
		log.Panic(err)
	}
	db.Create(&event)
	return event
}

func (repo EventRepository) GetEvent(id uuid.UUID) model.EventModel {
	db, err := gorm.Open("postgres", "user=postgres dbname=ta4 sslmode=disable password=docker")
	if err != nil {
		log.Panic(err)
	}

	var result = model.EventModel{Id: id}
	db.Find(&result, result)
	return result
}

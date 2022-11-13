package repository

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"log"
)

type EventModel struct {
	Id          uuid.UUID `gorm:"type:uuid;primary_key"`
	Sku         string    `gorm:"size:255;column:sku"`
	Description string    `gorm:"size:255;column:description"`
	ProjectId   uuid.UUID `gorm:"type:uuid;column:project_id"`
}

func (EventModel) TableName() string {
	return "Event"
}

type EventRepository struct{}

func (repo EventRepository) CreateEvent(event EventModel) EventModel {
	db, err := gorm.Open("postgres", "user=postgres dbname=ta4 sslmode=disable password=docker")
	if err != nil {
		log.Panic(err)
	}
	db.Create(&event)
	return event
}

func (repo EventRepository) GetEvent(id uuid.UUID) EventModel {
	db, err := gorm.Open("postgres", "user=postgres dbname=ta4 sslmode=disable password=docker")
	if err != nil {
		log.Panic(err)
	}

	var result = EventModel{Id: id}
	db.Find(&result, result)
	return result
}

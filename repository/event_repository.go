package repository

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"ta4/mod/model"
)

type EventRepository struct {
	store *gorm.DB
}

func NewEventRepository(store *gorm.DB) EventRepository {
	return EventRepository{store: store}
}

func (repo EventRepository) CreateEvent(event model.EventModel) model.EventModel {
	repo.store.Create(&event)
	return event
}

func (repo EventRepository) GetEvent(id uuid.UUID) model.EventModel {
	var result = model.EventModel{Id: id}
	repo.store.Find(&result, result)
	return result
}

package repository

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"ta4/mod/model"
)

type TriggerRepository struct {
	store *gorm.DB
}

func NewTriggerRepository(store *gorm.DB) TriggerRepository {
	return TriggerRepository{store: store}
}

func (repo TriggerRepository) CreateTrigger(Trigger model.TriggerModel) model.TriggerModel {
	repo.store.Create(&Trigger)
	return Trigger
}

func (repo TriggerRepository) CreateTriggerToEventRelation(Trigger model.TriggerToEventModel) model.TriggerToEventModel {
	repo.store.Create(&Trigger)
	return Trigger
}

func (repo TriggerRepository) GetTrigger(id uuid.UUID) model.TriggerModel {
	var result = model.TriggerModel{Id: id}
	repo.store.Find(&result, result)
	return result
}

package model

import "github.com/google/uuid"

type EventModel struct {
	Id          uuid.UUID `gorm:"type:uuid;primary_key"`
	Sku         string    `gorm:"size:255;column:sku"`
	Description string    `gorm:"size:255;column:description"`
	ProjectId   uuid.UUID `gorm:"type:uuid;column:project_id"`
}

func (EventModel) TableName() string {
	return "Event"
}

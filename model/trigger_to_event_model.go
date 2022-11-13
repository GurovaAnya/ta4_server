package model

import "github.com/google/uuid"

type TriggerToEventModel struct {
	Id         uuid.UUID `gorm:"type:uuid;primary_key"`
	EventId    uuid.UUID `gorm:"type:uuid;column:event_id"`
	TriggerId  uuid.UUID `gorm:"type:uuid;column:trigger_id"`
	ProjectId  uuid.UUID `gorm:"type:uuid;column:project_id"`
	EventCount int64     `gorm:"type:uuid;column:event_count"`
}

func (TriggerToEventModel) TableName() string {
	return "trigger_event"
}

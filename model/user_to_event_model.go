package model

import "github.com/google/uuid"

type UserToEventModel struct {
	Id         uuid.UUID `gorm:"type:uuid;primary_key"`
	EventId    uuid.UUID `gorm:"type:uuid;column:event_id"`
	UserId     uuid.UUID `gorm:"type:uuid;column:user_id"`
	ProjectId  uuid.UUID `gorm:"type:uuid;column:project_id"`
	IsRedeemed bool      `gorm:"column:is_redeemed"`
}

func (UserToEventModel) TableName() string {
	return "user_event"
}

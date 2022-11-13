package model

import "github.com/google/uuid"

type TriggerModel struct {
	Id            uuid.UUID `gorm:"type:uuid;primary_key"`
	Sku           string    `gorm:"size:255;column:sku"`
	Description   string    `gorm:"size:255;column:description"`
	AchievementId uuid.UUID `gorm:"type:uuid;column:achievement_id"`
	ProjectId     uuid.UUID `gorm:"type:uuid;column:project_id"`
}

func (TriggerModel) TableName() string {
	return "trigger"
}

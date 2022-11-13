package model

import "github.com/google/uuid"

type AchievementModel struct {
	Id          uuid.UUID `gorm:"type:uuid;primary_key"`
	Sku         string    `gorm:"size:255;column:sku"`
	Description string    `gorm:"size:255;column:description"`
	ImageUrl    string    `gorm:"size:255;column:image"`
	ProjectId   uuid.UUID `gorm:"type:uuid;column:project_id"`
}

func (AchievementModel) TableName() string {
	return "achievement"
}

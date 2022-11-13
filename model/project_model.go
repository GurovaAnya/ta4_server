package model

import "github.com/google/uuid"

type ProjectModel struct {
	Id         uuid.UUID `gorm:"type:uuid;primary_key"`
	WebhookUrl string    `gorm:"size:255;column:webhook_url"`
	Name       string    `gorm:"size:255;column:name"`
	ApiKey     string    `gorm:"size:255:column:api_key"`
}

func (ProjectModel) TableName() string {
	return "project"
}

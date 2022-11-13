package repository

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"log"
)

type ProjectModel struct {
	Id         uuid.UUID `gorm:"type:uuid;primary_key"`
	WebhookUrl string    `gorm:"size:255;column:webhook_url"`
	Name       string    `gorm:"size:255;column:name"`
	ApiKey     string    `gorm:"size:255:column:api_key"`
}

func (ProjectModel) TableName() string {
	return "project"
}

type ProjectRepository struct{}

func (repo ProjectRepository) GetProject(apiKey string) ProjectModel {
	db, err := gorm.Open("postgres", "user=postgres dbname=ta4 sslmode=disable password=docker")
	if err != nil {
		log.Panic(err)
	}

	var result = ProjectModel{ApiKey: apiKey}
	db.Find(&result, result)
	return result
}

func (repo ProjectRepository) CreateProject(project ProjectModel) ProjectModel {
	db, err := gorm.Open("postgres", "user=postgres dbname=ta4 sslmode=disable password=docker")
	if err != nil {
		log.Panic(err)
	}
	db.Create(&project)
	return project
}

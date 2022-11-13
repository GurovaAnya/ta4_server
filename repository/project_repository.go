package repository

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"log"
	"ta4/mod/model"
)

type ProjectRepository struct{}

func (repo ProjectRepository) GetProject(apiKey string) model.ProjectModel {
	db, err := gorm.Open("postgres", "user=postgres dbname=ta4 sslmode=disable password=docker")
	if err != nil {
		log.Panic(err)
	}

	var result = model.ProjectModel{ApiKey: apiKey}
	db.Find(&result, result)
	return result
}

func (repo ProjectRepository) CreateProject(project model.ProjectModel) model.ProjectModel {
	db, err := gorm.Open("postgres", "user=postgres dbname=ta4 sslmode=disable password=docker")
	if err != nil {
		log.Panic(err)
	}
	db.Create(&project)
	return project
}

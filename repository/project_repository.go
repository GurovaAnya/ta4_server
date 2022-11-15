package repository

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"ta4/mod/model"
)

type ProjectRepository struct {
	store *gorm.DB
}

func NewProjectRepository(store *gorm.DB) ProjectRepository {
	return ProjectRepository{store: store}
}

func (repo *ProjectRepository) GetProject(apiKey string) model.ProjectModel {
	db := repo.store
	var result = model.ProjectModel{ApiKey: apiKey}
	db.Find(&result, result)
	return result
}

func (repo *ProjectRepository) CreateProject(project model.ProjectModel) model.ProjectModel {
	db := repo.store
	db.Create(&project)
	return project
}

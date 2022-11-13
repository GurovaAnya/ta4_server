package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"log"
	"ta4/mod/internal/server/models"
)

type ProjectRepository interface {
	GetProject(ctx context.Context, id string) (*models.Project, error)
}

type ProjectModel struct {
	Id         uuid.UUID `gorm:"type:uuid;primary_key";"AUTO_INCREMENT"`
	WebhookUrl string    `gorm:"size:255"`
	Name       string    `gorm:"size:255"`
	ApiKey     string    `gorm:"size:255"`
}

func (ProjectModel) TableName() string {
	return "project"
}

type PostgresRepository struct{}

func (repo PostgresRepository) GetProject(id uuid.UUID) ProjectModel {
	db, err := gorm.Open("postgres", "user=postgres dbname=ta4 sslmode=disable password=docker")
	if err != nil {
		log.Panic(err)
	}

	var result ProjectModel
	db.Model(ProjectModel{Id: id}).Find(&result)
	return result
}

func (repo PostgresRepository) CreateProject(project ProjectModel) {

}

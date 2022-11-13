package repository

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"log"
)

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

type AchievementRepository struct{}

func (repo AchievementRepository) CreateAchievement(achievement AchievementModel) AchievementModel {
	db, err := gorm.Open("postgres", "user=postgres dbname=ta4 sslmode=disable password=docker")
	if err != nil {
		log.Panic(err)
	}
	db.Create(&achievement)
	return achievement
}

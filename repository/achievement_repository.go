package repository

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"log"
	"ta4/mod/model"
)

type AchievementRepository struct{}

func (repo AchievementRepository) CreateAchievement(achievement model.AchievementModel) model.AchievementModel {
	db, err := gorm.Open("postgres", "user=postgres dbname=ta4 sslmode=disable password=docker")
	if err != nil {
		log.Panic(err)
	}
	db.Create(&achievement)
	return achievement
}

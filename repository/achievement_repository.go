package repository

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"ta4/mod/model"
)

type AchievementRepository struct {
	store *gorm.DB
}

func NewAchievementRepository(store *gorm.DB) AchievementRepository {
	return AchievementRepository{store: store}
}

func (repo AchievementRepository) CreateAchievement(achievement model.AchievementModel) model.AchievementModel {
	repo.store.Create(&achievement)
	return achievement
}

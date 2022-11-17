package model

import "github.com/google/uuid"

type UserToAchievementModel struct {
	Id            uuid.UUID `gorm:"type:uuid;primary_key"`
	AchievementId uuid.UUID `gorm:"type:uuid;column:achievement_id"`
	UserId        uuid.UUID `gorm:"type:uuid;column:trigger_id"`
	ProjectId     uuid.UUID `gorm:"type:uuid;column:project_id"`
}

func (UserToAchievementModel) TableName() string {
	return "user_achievement"
}

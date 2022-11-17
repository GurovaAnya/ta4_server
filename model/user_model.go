package model

import "github.com/google/uuid"

type UserModel struct {
	Id            uuid.UUID `gorm:"type:uuid;primary_key"`
	WalletAddress string    `gorm:"size:255;column:wallet_address"`
	ProjectId     uuid.UUID `gorm:"type:uuid;column:project_id"`
}

func (UserModel) TableName() string {
	return "user"
}

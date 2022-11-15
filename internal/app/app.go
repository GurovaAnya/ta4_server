package app

import (
	"github.com/jinzhu/gorm"
	"ta4/mod/config"
)

type Application struct {
	Store  *gorm.DB
	Config config.Config
}

func NewApplication(
	storeInstance *gorm.DB,
	configInstance *config.Config,
) Application {
	return Application{
		Store:  storeInstance,
		Config: *configInstance,
	}
}

package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	ServerHost       string `env:"SERVER_HOST" env-required:"true" env-description:"Host of server instance"`
	ServerPort       int    `env:"SERVER_PORT" env-required:"true" env-description:"Port of server instance"`
	PostgresHost     string `env:"DB_HOST" env-required:"true" env-description:"Host of Postgres db instance"`
	PostgresPort     int    `env:"DB_PORT" env-required:"true" env-description:"Port of Postgres instance"`
	PostgresUser     string `env:"DB_USER" env-required:"true" env-description:"User of Postgres instance"`
	PostgresPassword string `env:"DB_PASSWORD" env-required:"true" env-description:"Password of Postgres instance"`
	PostgresDbName   string `env:"DB_NAME" env-required:"true" env-description:"Db name of Postgres instance"`
}

func GetConfigDescription() (string, error) {
	return cleanenv.GetDescription(&Config{}, nil)
}

func Init() (Config, error) {
	c := Config{}
	err := cleanenv.ReadEnv(&c)
	return c, err
}

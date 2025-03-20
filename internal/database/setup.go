package database

import (
	"app/internal/config"
	"app/internal/entity"
	"errors"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectAndMigrate(cfg *config.Config) (*gorm.DB, error) {
	database, err := gorm.Open(postgres.Open(cfg.Database.DSN))
	if err != nil {
		return nil, errors.New("failed connect to database")
	}

	err = database.AutoMigrate(entity.User{})

	if err != nil {
		return nil, errors.New("failed auto migrate database")
	}

	return database, nil
}

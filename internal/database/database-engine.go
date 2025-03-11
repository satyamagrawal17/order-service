package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"ordering_service/internal/model"
)

type DatabaseEngine struct {
	Db *gorm.DB
}

func InitDatabaseEngine() (*DatabaseEngine, error) {
	db, err := CreateDatabaseEngine()
	if err != nil {
		return &DatabaseEngine{}, err
	}

	return &DatabaseEngine{
		Db: db,
	}, err
}

func CreateDatabaseEngine() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(
		&model.Order{},
		&model.OrderItem{},
	)

	return db, nil
}

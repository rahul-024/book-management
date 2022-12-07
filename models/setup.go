package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "postgres://postgres:root@localhost:5432/gopractise"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	err = database.AutoMigrate(&Book{})
	if err != nil {
		return
	}
	DB = database
}

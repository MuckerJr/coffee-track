package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Coffee struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

type Recipe struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `json:"name"`
	Ingredients string `json:"ingredients"`
}

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("coffee.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB.AutoMigrate(&Coffee{}, &Recipe{})
}

package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Coffee struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"not null"`
	Size     string `gorm:"not null"`
	Quantity string `gorm:"not null"`
	Vendor   string
	Roast    string `gorm:"not null"`
	Grind    string `gorm:"not null"`
}

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("coffee.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}

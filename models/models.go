package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Roast int

const (
	Light Roast = iota + 1
	Medium
	Dark
)

func (r Roast) String() string {
	switch r {
	case Light:
		return "Light"
	case Medium:
		return "Medium"
	case Dark:
		return "Dark"
	default:
		return "Unknown"
	}
}

type Grind int

const (
	Wholebean Grind = iota + 1
	Ground
)

func (g Grind) String() string {
	switch g {
	case Wholebean:
		return "Wholebean"
	case Ground:
		return "Ground"
	default:
		return "Unknown"
	}
}

type Coffee struct {
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	Name     string `gorm:"not null"`
	Size     string `gorm:"not null"`
	Quantity string `gorm:"not null"`
	Vendor   string `gorm:"not null"`
	Roast    Roast  `gorm:"not null"`
	Grind    Grind  `gorm:"not null"`
}

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("coffee.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	if err := DB.AutoMigrate(&Coffee{}); err != nil {
		panic("failed to migrate database")
	}
}

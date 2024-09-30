package models

import (
	"encoding/json"
	"errors"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Roast int

const (
	LightRoast Roast = iota + 1
	MediumRoast
	DarkRoast
)

func (r Roast) String() string {
	switch r {
	case LightRoast:
		return "Light"
	case MediumRoast:
		return "Medium"
	case DarkRoast:
		return "Dark"
	default:
		return "Unknown"
	}
}

// Custom JSON marshalling for Roast
func (r Roast) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

// Custom JSON unmarshalling for Roast
func (r *Roast) UnmarshalJSON(data []byte) error {
	var roast string
	if err := json.Unmarshal(data, &roast); err != nil {
		return err
	}
	switch roast {
	case "Light":
		*r = LightRoast
	case "Medium":
		*r = MediumRoast
	case "Dark":
		*r = DarkRoast
	default:
		return errors.New("invalid roast")
	}
	return nil
}

type Grind int

const (
	Wholebean Grind = iota + 1
	Ground
	Fine
	Medium
	Coarse
)

func (g Grind) String() string {
	switch g {
	case Wholebean:
		return "Wholebean"
	case Ground:
		return "Ground"
	case Fine:
		return "Fine"
	case Medium:
		return "Medium"
	case Coarse:
		return "Coarse"
	default:
		return "Unknown"
	}
}

// Custom JSON marshalling for Grind
func (g Grind) MarshalJSON() ([]byte, error) {
	return json.Marshal(g.String())
}

// Custom JSON unmarshalling for Grind
func (g *Grind) UnmarshalJSON(data []byte) error {
	var grind string
	if err := json.Unmarshal(data, &grind); err != nil {
		return err
	}
	switch grind {
	case "Wholebean":
		*g = Wholebean
	case "Ground":
		*g = Ground
	case "Fine":
		*g = Fine
	case "Medium":
		*g = Medium
	case "Coarse":
		*g = Coarse
	default:
		return errors.New("invalid grind")
	}
	return nil
}

type Coffee struct {
	ID       uint           `gorm:"primaryKey;autoIncrement"`
	Name     string         `gorm:"not null"`
	Vendor   string         `gorm:"not null"`
	Quantity int            `gorm:"not null;default:0"`
	Details  []CoffeeDetail `gorm:"foreignKey:CoffeeID"`
}

type CoffeeDetail struct {
	ID       uint     `gorm:"primaryKey;autoIncrement"`
	CoffeeID uint     `gorm:"not null"`
	Size     []string `gorm:"not null"`
	Grind    Grind    `gorm:"not null"`
	Roast    Roast    `gorm:"not null"`
}

type Recipe struct {
	ID         uint   `gorm:"primaryKey;autoIncrement"`
	CoffeeUsed uint   `gorm:"not null"`
	WaterUsed  uint   `gorm:"not null"`
	BrewMethod string `gorm:"not null"`
	GrindSize  Grind  `gorm:"not null"`
}

type Brew struct {
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	CoffeeID uint   `gorm:"not null"`
	RecipeID uint   `gorm:"not null"`
	Rating   uint   `gorm:"not null"`
	Notes    string `gorm:"not null"`
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

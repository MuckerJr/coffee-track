package cli

import (
	"fmt"
	"os"

	"github.com/charmbracelet/huh"

	"coffee-track/models"
)

// type Coffee struct {
// 	name     string
// 	size     string
// 	quantity string
// 	vendor   string
// 	roast    Roast
// 	grind    Grind
// }

// type Roast int

// const (
// 	Light Roast = iota + 1
// 	Medium
// 	Dark
// )

// func (r Roast) String() string {
// 	switch r {
// 	case Light:
// 		return "Light"
// 	case Medium:
// 		return "Medium"
// 	case Dark:
// 		return "Dark"
// 	default:
// 		return "Unknown"
// 	}
// }

// type Grind int

// const (
// 	Wholebean Grind = iota + 1
// 	Ground
// )

// func (g Grind) String() string {
// 	switch g {
// 	case Wholebean:
// 		return "Wholebean"
// 	case Ground:
// 		return "Ground"
// 	default:
// 		return "Unknown"
// 	}
// }

func runAddCoffeeForm() {
	var coffee models.Coffee
	form := huh.NewForm(
		huh.NewGroup(huh.NewNote().Title("Add Coffee").Description("Add a new coffee to your inventory").Next(true).NextLabel("Next")),
		huh.NewGroup(
			huh.NewInput().Title("Name").Prompt("Enter coffee name").Validate(ValidateNonEmpty),
			huh.NewInput().Title("Size").Prompt("Amount of coffee in grams").Validate(ValidateNonEmpty),
			huh.NewInput().Title("Quantity").Prompt("Number of bags").Validate(ValidateNonEmpty),
			huh.NewInput().Title("Vendor").Prompt("From whom did you purchase this coffee").Validate(ValidateNonEmpty),
			huh.NewSelect[models.Roast]().Title("Roast").Options(
				huh.NewOption("Light", models.Light).Selected(true),
				huh.NewOption("Medium", models.Medium),
				huh.NewOption("Dark", models.Dark),
			),
			huh.NewSelect[models.Grind]().Title("Grind").Options(
				huh.NewOption("Wholebean", models.Wholebean).Selected(true),
				huh.NewOption("Ground", models.Ground),
			),
		),
	)

	if err := form.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	addCoffee(&coffee)
}

func addCoffee(coffee *models.Coffee) {
	fmt.Println("Adding coffee...")

	models.InitDB()

	if err := models.DB.Create(coffee).Error; err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	} else {
		fmt.Println("Coffee Added!")
	}
}

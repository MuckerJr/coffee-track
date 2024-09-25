package cli

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"coffee-track/models"
)

type formModel struct {
	fieldIndex int
	fields     []string
	values     map[string]string
	quitting   bool
	submitted  bool
}

func intitialFormModel() formModel {
	fields := []string{"Name", "Size (grams)", "Quantity", "Vendor (optional)", "Roast (light, medium, dark)", "Grind (whole-bean, ground)"}
	values := make(map[string]string)
	for _, f := range fields {
		values[f] = ""
	}
	return formModel{fields: fields, values: values}
}

func (m formModel) Init() tea.Cmd {
	return nil
}

func (m formModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit
		case "enter":
			if m.fieldIndex < len(m.fields)-1 {
				m.fieldIndex++
			} else {
				m.submitted = true
				return m, tea.Quit
			}
		case "backspace":
			if len(m.values[m.fields[m.fieldIndex]]) > 0 {
				m.values[m.fields[m.fieldIndex]] = m.values[m.fields[m.fieldIndex]][:len(m.values[m.fields[m.fieldIndex]])-1]
			}
		default:
			m.values[m.fields[m.fieldIndex]] += msg.String()
		}
	}

	return m, nil
}

func (m formModel) View() string {
	if m.quitting {
		return "Goodbye!"
	}
	if m.submitted {
		return "Coffee Submitted!\n"
	}

	var b strings.Builder
	b.WriteString("Add Coffee\n\n")
	for i, field := range m.fields {
		if i == m.fieldIndex {
			b.WriteString(fmt.Sprintf("> %s: %s\n", field, m.values[field]))
		} else {
			b.WriteString(fmt.Sprintf("  %s: %s\n", field, m.values[field]))
		}
	}
	return b.String()
}

func runAddCoffeeForm() {
	p := tea.NewProgram(intitialFormModel())
	m, err := p.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	if m, ok := m.(formModel); ok && m.submitted {
		addCoffee(m.values)
	} else {
		fmt.Println("Form not submitted")
	}
}

func addCoffee(values map[string]string) {
	fmt.Println("Adding coffee...")
	db, err := gorm.Open(sqlite.Open("coffee.db"), &gorm.Config{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	coffee := models.Coffee{
		Name:     values["Name"],
		Size:     values["Size (grams)"],
		Quantity: values["Quantity"],
		Vendor:   values["Vendor (optional)"],
		Roast:    values["Roast (light, medium, dark)"],
		Grind:    values["Grind (whole-bean, ground)"],
	}

	if err := db.Create(&coffee).Error; err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Coffee added!")
}

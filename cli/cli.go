package cli

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

type cliModel struct {
	list      list.Model
	quitting  bool
	confirmed bool
	selected  string
}

func initialCLIModel() cliModel {
	items := []list.Item{
		item{title: "Add Coffee", desc: "Add coffee to the inventory"},
		item{title: "List Coffees", desc: "List all coffees in the inventory"},
	}

	const defaultWidth = 50
	const listHeight = 20

	l := list.New(items, list.NewDefaultDelegate(), defaultWidth, listHeight)
	l.Title = "Select Action"

	return cliModel{list: l}
}

func (m cliModel) Init() tea.Cmd {
	return nil
}

func (m cliModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit
		case "enter":
			i, ok := m.list.SelectedItem().(item)
			if ok {
				m.confirmed = true
				m.selected = i.title
				return m, tea.Quit
			}
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m cliModel) View() string {
	if m.quitting {
		return "Goodbye!\n"
	} else {
		return "\n" + m.list.View()
	}
}

func RunCLI() {
	for {
		p := tea.NewProgram(initialCLIModel())
		m, err := p.Run()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		if m, ok := m.(cliModel); ok && m.confirmed {
			switch m.selected {
			case "Add Coffee":
				runAddCoffeeForm()
			case "List Coffees":
				listCoffees()
			default:
				fmt.Println("Invalid selection")
			}
		} else {
			fmt.Println("Exiting...")
			break
		}
	}
}

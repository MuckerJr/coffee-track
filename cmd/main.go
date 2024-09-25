package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"

	"coffee-track/cli"
	"coffee-track/server"
)

type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

type mainModel struct {
	list     list.Model
	choice   string
	quitting bool
}

func initialMainModel() mainModel {
	items := []list.Item{
		item{title: "Start Server", desc: "Run the web server"},
		item{title: "Start CLI", desc: "Run the command-line interface"},
	}

	const defaultWidth = 20
	const listHeight = 14

	l := list.New(items, list.NewDefaultDelegate(), defaultWidth, listHeight)
	l.Title = "Select Mode"

	return mainModel{list: l}
}

func (m mainModel) Init() tea.Cmd {
	return nil
}

func (m mainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit
		case "enter":
			i, ok := m.list.SelectedItem().(item)
			if ok {
				m.choice = i.title
				return m, tea.Quit
			}
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m mainModel) View() string {
	if m.choice != "" {
		return fmt.Sprintf("You chose: %s\n", m.choice)
	}
	if m.quitting {
		return "Goodbye!\n"
	}
	return "\n" + m.list.View()
}

func main() {
	p := tea.NewProgram(initialMainModel())
	m, err := p.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	if m, ok := m.(mainModel); ok {
		switch m.choice {
		case "Start Server":
			server.RunServer()
		case "Start CLI":
			cli.RunCLI()
		default:
			fmt.Println("Closing...")
		}
	} else {
		fmt.Println("Failed to get the model")
	}

}

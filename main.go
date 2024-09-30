package main

import (
	"fmt"
	"os"

	"github.com/Piitschy/twaskwarrior-tui/internal/tw"
	"github.com/Piitschy/twaskwarrior-tui/views"
	tea "github.com/charmbracelet/bubbletea"
)

type sessionState int

const (
	tasktableView sessionState = iota
	newTaskView
	contextView
)

type MainModel struct {
	state     sessionState
	tasktable tea.Model
}

func InitModel(tw *tw.TaskWarrior, columns []string) MainModel {
	tasktableView := views.InitTasktableView(tw, columns)
	return MainModel{
		tasktable: tasktableView,
	}
}

func (m MainModel) Init() tea.Cmd {
	return nil
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.tasktable, cmd = m.tasktable.Update(msg)
	return m, cmd
}

func (m MainModel) View() string {
	return m.tasktable.View()
}

func main() {
	tw, err := tw.NewTaskWarrior()
	if err != nil {
		panic(err)
	}

	tw.AddFilter("status", "pending")

	columns := []string{"ID", "Project", "Description", "Status"}

	m := InitModel(tw, columns)
	p := tea.NewProgram(m)
	p.EnterAltScreen()
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

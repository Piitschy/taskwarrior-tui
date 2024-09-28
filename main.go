package main

import (
	"fmt"
	"os"

	"github.com/Piitschy/twaskwarrior-tui/internal/tw"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	tw, err := tw.NewTaskWarrior()
	if err != nil {
		panic(err)
	}

	columns := []string{"ID", "Project", "Description", "Status"}

	m := InitModel(tw, columns)
	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

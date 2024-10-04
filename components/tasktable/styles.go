package tasktable

import "github.com/charmbracelet/lipgloss"

var HeaderStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#FFA500")).
	Bold(true)

var RowStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#FFFFFF"))

var SelectedRowStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#000000")).
	Background(lipgloss.Color("#FFFFFF"))

// background orange
var ActiveRowStyle = lipgloss.NewStyle().
	Background(lipgloss.Color("#FFA500")).
	Foreground(lipgloss.Color("#FFFFFF"))

var NextRowStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#fffa00"))

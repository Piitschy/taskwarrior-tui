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

var SelectedActiveRowStyle = lipgloss.NewStyle().
	Background(lipgloss.Color("#FFFFFF")).
	Foreground(lipgloss.Color("#FFA500"))

var NextRowStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#FFFA00"))

var SelectedNextRowStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#000000")).
	Background(lipgloss.Color("#FFFA50"))

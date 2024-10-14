package activefilters

import "github.com/charmbracelet/lipgloss"

var RowStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#FFFFFF"))

var SelectedRowStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#000000")).
	Background(lipgloss.Color("#FFFFFF"))

var DisabledRowStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#808080"))

var SelectedDisabledRowStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#808080")).
	Background(lipgloss.Color("#FFFFFF"))

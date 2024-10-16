package views

import (
	"github.com/Piitschy/twaskwarrior-tui/components/activefilters"
	"github.com/Piitschy/twaskwarrior-tui/components/tasktable"
	"github.com/Piitschy/twaskwarrior-tui/internal/tw"
	"github.com/Piitschy/twaskwarrior-tui/internal/utils"
	"github.com/Piitschy/twaskwarrior-tui/keymap"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type sessionState int

const (
	none = iota
	showNewFilter
	showFilters
)

type TasktableView struct {
	tw          *tw.TaskWarrior
	tasktable   tea.Model
	state       sessionState
	filterInput textinput.Model
	filterList  tea.Model
	width       int
}

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("56"))

var inactiveStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("24"))

func InitTasktableView(tw *tw.TaskWarrior, columns []string, expandedColumn int) TasktableView {
	// table
	tasktable := tasktable.InitModel(tw, columns, expandedColumn)
	// filter
	filterInput := textinput.New()
	filterInput.Placeholder = "add filter like: 'project:work'"
	filterInput.ShowSuggestions = true
	filterInput.KeyMap.AcceptSuggestion = keymap.KeyMap.AcceptSuggestion
	filterInput.SetSuggestions(utils.ProjectSuggestions(tw.GetProjects()))
	// help
	return TasktableView{
		state:       none,
		tw:          tw,
		tasktable:   tasktable,
		filterInput: filterInput,
		filterList:  activefilters.InitModel(tw),
	}
}

func (m TasktableView) Init() tea.Cmd {
	return nil
}

func (m TasktableView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch {
	case m.state == showNewFilter:
		m.filterInput, cmd = m.filterInput.Update(msg)
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch {
			case msg.String() == "esc":
				m.filterInput.Blur()
				m.state = none
				utils.BlockCommentLine = false
			case msg.String() == "enter":
				err := m.tw.AddFilterFromString(m.filterInput.Value())
				if err == nil {
					m.filterInput.SetValue("")
					m.filterInput.Blur()
					m.state = none
					utils.BlockCommentLine = false
				}
			}
		}
		return m, cmd
	case m.state == showFilters:
		m.filterList, cmd = m.filterList.Update(msg)
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch {
			case key.Matches(msg, keymap.KeyMap.Filter):
				m.state = showNewFilter
				m.filterInput.Focus()
				utils.BlockCommentLine = true
			case key.Matches(msg, keymap.KeyMap.Quit) || key.Matches(msg, keymap.KeyMap.Up) || key.Matches(msg, keymap.KeyMap.Down):
				m.state = none
				utils.BlockCommentLine = false
			}
		}
		return m, cmd
	}
	m.tasktable, cmd = m.tasktable.Update(msg)
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keymap.KeyMap.Filter):
			m.state = showNewFilter
			m.filterInput.Focus()
			utils.BlockCommentLine = true
		case key.Matches(msg, keymap.KeyMap.Left) || key.Matches(msg, keymap.KeyMap.Right):
			m.state = showFilters
		}
	}
	return m, cmd
}

func (m TasktableView) View() string {
	tasktableView := m.tasktable.View()
	filterView := ""
	if m.state == none {
		filterView += inactiveStyle.Width(m.width - 2).Render(m.filterList.View() + "\n\n" + m.filterInput.View())
	} else {
		filterView += baseStyle.Width(m.width - 2).Render(m.filterList.View() + "\n\n" + m.filterInput.View())
	}
	return tasktableView + "\n" + filterView
}

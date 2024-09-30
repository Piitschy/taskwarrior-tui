package views

import (
	"fmt"

	"github.com/Piitschy/twaskwarrior-tui/components/activefilters"
	"github.com/Piitschy/twaskwarrior-tui/components/tasktable"
	"github.com/Piitschy/twaskwarrior-tui/internal/tw"
	"github.com/Piitschy/twaskwarrior-tui/keymap"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
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
	help        help.Model
}

func InitTasktableView(tw *tw.TaskWarrior, columns []string) TasktableView {
	// table
	tasktable := tasktable.InitModel(tw, columns)
	// filter
	filterInput := textinput.New()
	filterInput.Placeholder = "add filter like: 'project:work'"
	// help
	help := help.New()
	return TasktableView{
		state:       none,
		tw:          tw,
		tasktable:   tasktable,
		filterInput: filterInput,
		filterList:  activefilters.InitModel(tw),
		help:        help,
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
			case msg.String() == "esc" || key.Matches(msg, keymap.KeyMap.Filter):
				m.filterInput.Blur()
				m.state = none
			case msg.String() == "enter":
				m.tw.AddFilterFromString(m.filterInput.Value())
				m.filterInput.Blur()
				m.filterInput.SetValue("")
				m.state = none

			}
		}
		return m, cmd
	case m.state == showFilters:
		m.filterList, cmd = m.filterList.Update(msg)
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch {
			case key.Matches(msg, keymap.KeyMap.Quit):
				m.state = none
			}
		}
		return m, cmd
	}
	m.tasktable, cmd = m.tasktable.Update(msg)
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keymap.KeyMap.Help):
			m.help.ShowAll = !m.help.ShowAll
		case key.Matches(msg, keymap.KeyMap.Filter):
			m.state = showNewFilter
			m.filterInput.Focus()
		case key.Matches(msg, keymap.KeyMap.ActiveFilters):
			m.state = showFilters
		}
	}
	return m, cmd
}

func (m TasktableView) View() string {
	tasktableView := m.tasktable.View()
	var filterView string
	switch m.state {
	case showNewFilter:
		filterView = m.filterInput.View()
	case showFilters:
		filterView = m.filterList.View()
	}
	helpView := m.help.View(keymap.KeyMap)
	return fmt.Sprintf("%s\n%s\n\n%s", tasktableView, filterView, helpView)
}

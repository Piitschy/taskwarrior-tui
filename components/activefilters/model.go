package activefilters

import (
	"github.com/Piitschy/twaskwarrior-tui/internal/tw"
	"github.com/Piitschy/twaskwarrior-tui/keymap"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	tw     *tw.TaskWarrior
	cursor int
}

func InitModel(tw *tw.TaskWarrior) model {
	return model{tw: tw, cursor: 0}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keymap.KeyMap.Right):
			if m.cursor < len(m.tw.GetFilters())-1 {
				m.cursor++
			}
		case key.Matches(msg, keymap.KeyMap.Left):
			if m.cursor > 0 {
				m.cursor--
			}
		case key.Matches(msg, keymap.KeyMap.DisableFilter):
			selectedFilter := m.tw.GetFilters()[m.cursor]
			m.tw.ToggleDisableFilter(selectedFilter)
		case key.Matches(msg, keymap.KeyMap.Delete):
			selectedFilter := m.tw.GetFilters()[m.cursor]
			m.tw.RemoveFilter(selectedFilter)
			m.cursor--
		}
	}
	return m, cmd
}

func (m model) View() string {
	s := ""
	if len(m.tw.GetFilters()) == 0 {
		return "No active filters"
	}
	for i, filter := range m.tw.GetFilters() {
		if !filter.Disabled {
			if i == m.cursor {
				s += SelectedRowStyle.Render(filter.String())
			} else {
				s += RowStyle.Render(filter.String())
			}
		} else {
			if i == m.cursor {
				s += SelectedDisabledRowStyle.Render(filter.String())
			} else {
				s += DisabledRowStyle.Render(filter.String())
			}
		}
		s += " + "
	}
	s = s[:len(s)-3]
	return s
}

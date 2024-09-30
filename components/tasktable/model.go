package tasktable

import (
	"reflect"
	"slices"
	"strconv"

	"github.com/Piitschy/twaskwarrior-tui/internal/tw"
	"github.com/Piitschy/twaskwarrior-tui/internal/utils"
	"github.com/Piitschy/twaskwarrior-tui/keymap"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

type Model struct {
	table   table.Table
	columns []string
	tw      *tw.TaskWarrior
	cursor  int
}

func InitModel(tw *tw.TaskWarrior, columns []string) Model {
	t := table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("99"))).
		BorderTop(true).BorderBottom(true).BorderLeft(true).BorderRight(true).BorderColumn(false).
		Headers(utils.SpaceAround(columns)...)
	m := Model{tw: tw, table: *t, cursor: 0, columns: columns}
	rows := m.getRows()
	m.table.Rows(rows...)
	return m
}

func (m Model) Init() tea.Cmd { return nil }

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {

		case key.Matches(msg, keymap.KeyMap.Down):
			if m.cursor < len(m.tw.GetFilteredTasks())-1 {
				m.cursor++
			}

		case key.Matches(msg, keymap.KeyMap.Up):
			if m.cursor > 0 {
				m.cursor--
			}

		case key.Matches(msg, keymap.KeyMap.Quit):
			return m, tea.Quit

		case key.Matches(msg, keymap.KeyMap.Space):
			taskId := m.tw.GetFilteredTasks()[m.cursor].Id
			activeTasks, _ := m.tw.GetActiveTasks()
			if activeTasks.Contains(taskId) {
				m.tw.StopTask(taskId)
			} else {
				m.tw.StartTask(taskId)
			}

		case key.Matches(msg, keymap.KeyMap.Done):
			taskId := m.tw.GetFilteredTasks()[m.cursor].Id
			m.tw.TaskDone(taskId)

		case key.Matches(msg, keymap.KeyMap.Undo):
			m.tw.Undo()

		}
	}
	m.tw.LoadTasks()
	sd := table.NewStringData(m.getRows()...)
	m.table.Data(sd)
	return m, cmd
}

func (m Model) View() string {
	activeTasks, _ := m.tw.GetActiveTasks()
	activeRows := []int{}
	for i, task := range m.tw.GetFilteredTasks() {
		if activeTasks.Contains(task.Id) {
			activeRows = append(activeRows, i+1)
		}
	}
	m.table.StyleFunc(func(row, col int) lipgloss.Style {
		switch {
		case row == 0:
			return HeaderStyle
		case slices.Contains(activeRows, row):
			return ActiveRowStyle
		case row == m.cursor+1:
			return SelectedRowStyle
		default:
			return RowStyle
		}
	})

	return m.table.Render()
}

func (m Model) getRows() [][]string {
	tasks := m.tw.GetFilteredTasks()
	rows := make([][]string, len(tasks))
	for i, task := range tasks {
		r := reflect.ValueOf(task)
		// create row by iterating over columns and getting the field ValueOf
		// id is a special case, because it is an int
		row := make([]string, len(m.columns))
		for i, fieldname := range m.columns {
			if i == 0 {
				row[i] = strconv.Itoa(task.Id)
				continue
			}
			row[i] = r.FieldByName(fieldname).String()
		}
		rows[i] = utils.SpaceAround(row)
	}
	return rows
}

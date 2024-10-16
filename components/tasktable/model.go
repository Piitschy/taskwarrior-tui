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
	table      table.Table
	columns    []string
	rows       [][]string
	activeRows []int
	nextRows   []int
	tw         *tw.TaskWarrior
	cursor     int
	width      int
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

func autoSpace(cols []string, width int, expandingCol int) []string {
	result := make([]string, len(cols))
	var totalWidth int
	for i, col := range cols {
		result[i] = " " + col + " "
		totalWidth += len(result[i])
	}
	if totalWidth >= width {
		return result
	}
	space := ""
	for i := 0; i < width-totalWidth-10; i++ {
		space += " "
	}
	result[expandingCol] = " " + cols[expandingCol] + space
	return result
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.table.
			Width(m.width).
			Headers(autoSpace(m.columns, m.width, 3)...)
	case tea.KeyMsg:
		var err error = nil
		switch {

		case key.Matches(msg, keymap.KeyMap.Down):
			if m.cursor < len(m.tw.GetFilteredTasks())-1 {
				m.cursor++
			} else {
				m.cursor = 0
			}

		case key.Matches(msg, keymap.KeyMap.Up):
			if m.cursor > 0 {
				m.cursor--
			} else {
				m.cursor = len(m.tw.GetFilteredTasks()) - 1
			}

		case key.Matches(msg, keymap.KeyMap.Quit):
			return m, tea.Quit

		case key.Matches(msg, keymap.KeyMap.Space):
			taskId := m.tw.GetFilteredTasks()[m.cursor].Id
			activeTasks, _ := m.tw.GetActiveTasks()
			if activeTasks.Contains(taskId) {
				err = m.tw.StopTask(taskId)
			} else {
				err = m.tw.StartTask(taskId)
			}

		case key.Matches(msg, keymap.KeyMap.Done):
			taskId := m.tw.GetFilteredTasks()[m.cursor].Id
			err = m.tw.TaskDone(taskId)

		case key.Matches(msg, keymap.KeyMap.Undo):
			err = m.tw.Undo()

		case key.Matches(msg, keymap.KeyMap.Next):
			taskId := m.tw.GetFilteredTasks()[m.cursor].Id
			nextTasks, nErr := m.tw.GetNextTasks()
			if nErr != nil {
				err = nErr
			}
			if nextTasks.Contains(taskId) {
				err = m.tw.TaskUnnext(taskId)
			} else {
				err = m.tw.TaskNext(taskId)
			}
		}
		if err != nil {
			panic(err)
		}
	}
	m.tw.LoadTasks()
	activeTasks, _ := m.tw.GetActiveTasks()
	nextTasks, _ := m.tw.GetNextTasks()
	m.activeRows = []int{}
	m.nextRows = []int{}
	for i, task := range m.tw.GetFilteredTasks() {
		if activeTasks.Contains(task.Id) {
			m.activeRows = append(m.activeRows, i+1)
		}
		if nextTasks.Contains(task.Id) {
			m.nextRows = append(m.nextRows, i+1)
		}
	}
	m.rows = m.getRows()
	sd := table.NewStringData(m.rows...)
	m.table.Data(sd)
	return m, cmd
}

func (m Model) View() string {
	m.table.StyleFunc(func(row, col int) lipgloss.Style {
		switch {
		case row == 0:
			return HeaderStyle
		case slices.Contains(m.activeRows, row) && row == m.cursor+1:
			return SelectedActiveRowStyle
		case slices.Contains(m.nextRows, row) && row == m.cursor+1:
			return SelectedNextRowStyle
		case row == m.cursor+1:
			return SelectedRowStyle
		case slices.Contains(m.activeRows, row) && slices.Contains(m.nextRows, row):
			return NextActiveRowStyle
		case slices.Contains(m.activeRows, row):
			return ActiveRowStyle
		case slices.Contains(m.nextRows, row):
			return NextRowStyle
		default:
			return RowStyle
		}
	})

	return m.table.Render()
}

func (m Model) getRows() [][]string {
	tasks := m.tw.GetFilteredTasks() // TODO: sort tasks
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
			t := r.FieldByName(fieldname).Type().Kind()
			var s string
			switch t {
			case reflect.String:
				s = r.FieldByName(fieldname).String()
			case reflect.Slice:
				for i := 0; i < r.FieldByName(fieldname).Len(); i++ {
					s += r.FieldByName(fieldname).Index(i).String() + " "
				}
			}

			row[i] = s
		}
		rows[i] = utils.SpaceAround(row)
	}
	return rows
}

package main

import (
	"reflect"
	"slices"
	"strconv"

	"github.com/Piitschy/twaskwarrior-tui/internal/tw"
	"github.com/Piitschy/twaskwarrior-tui/internal/utils"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

type model struct {
	table   table.Table
	columns []string
	tw      *tw.TaskWarrior
	cursor  int
}

func InitModel(tw *tw.TaskWarrior, columns []string) model {
	t := table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("99"))).
		BorderTop(true).BorderBottom(true).BorderLeft(true).BorderRight(true).BorderColumn(false).
		Headers(utils.SpaceAround(columns)...)
	m := model{tw: tw, table: *t, cursor: 0, columns: columns}
	rows := m.getRows()
	m.table.Rows(rows...)
	return m
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {

		case "j", "down":
			if m.cursor < len(m.tw.Tasks)-1 {
				m.cursor++
			}

		case "k", "up":
			if m.cursor > 0 {
				m.cursor--
			}

		case "q", "ctrl+c":
			return m, tea.Quit

		case " ":
			taskId := m.tw.Tasks[m.cursor].Id
			activeTasks, _ := m.tw.GetActiveTasks()
			if activeTasks.Contains(taskId) {
				m.tw.StopTask(taskId)
			} else {
				m.tw.StartTask(taskId)
			}

		case "d":
			taskId := m.tw.Tasks[m.cursor].Id
			m.tw.TaskDone(taskId)
		}
	}
	m.tw.LoadTasks()
	sd := table.NewStringData(m.getRows()...)
	m.table.Data(sd)
	return m, cmd
}

func (m model) View() string {
	activeTasks, _ := m.tw.GetActiveTasks()
	activeRows := []int{}
	for i, task := range m.tw.Tasks {
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

func (m model) getRows() [][]string {
	tasks := m.tw.Tasks
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

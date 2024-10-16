package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/Piitschy/twaskwarrior-tui/internal/tw"
	"github.com/Piitschy/twaskwarrior-tui/internal/utils"
	"github.com/Piitschy/twaskwarrior-tui/keymap"
	"github.com/Piitschy/twaskwarrior-tui/views"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type sessionState int

const (
	tasktableView sessionState = iota
	newTaskView
	contextView
)

type MainModel struct {
	state         sessionState
	tasktable     tea.Model
	activeCommand bool
	commandline   textinput.Model
	help          help.Model
	width         int
	height        int
}

func InitModel(tw *tw.TaskWarrior, columns []string) MainModel {
	tasktableView := views.InitTasktableView(tw, columns)
	cl := textinput.New()
	cl.Placeholder = "Enter taskwarrior command here..."
	suggestions := utils.AddProjectSuggestions(utils.Suggestions, tw.GetProjects())
	cl.ShowSuggestions = true
	cl.KeyMap.AcceptSuggestion = keymap.KeyMap.AcceptSuggestion
	cl.SetSuggestions(suggestions)
	cl.Prompt = ":task> "
	help := help.New()
	help.Width = 500
	return MainModel{
		tasktable:     tasktableView,
		activeCommand: false,
		commandline:   cl,
		help:          help,
		width:         500,
	}
}

func (m MainModel) Init() tea.Cmd {
	return nil
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyMsg:
		if m.activeCommand {
			switch {
			case key.Matches(msg, key.NewBinding(key.WithKeys("esc"))):
				m.commandline.Blur()
				m.activeCommand = false
			// case msg.String() == "tab":
			// 	m.commandline
			case msg.String() == "enter":
				command := strings.Split(m.commandline.Value(), " ")
				exec.Command("task", command...).Run()
				m.commandline.Blur()
				m.activeCommand = false
				m.commandline.SetValue("")
			}
			m.commandline, cmd = m.commandline.Update(msg)
			return m, cmd
		}
	}

	m.tasktable, cmd = m.tasktable.Update(msg)
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keymap.KeyMap.Command) && !utils.BlockCommentLine:
			m.activeCommand = !m.activeCommand
			m.commandline.Focus()
		case key.Matches(msg, keymap.KeyMap.Help):
			m.help.ShowAll = !m.help.ShowAll
		}
	}
	return m, cmd
}

func (m MainModel) View() string {
	helpView := m.help.View(keymap.KeyMap)
	view := fmt.Sprintf(
		"%s\n%s",
		m.tasktable.View(),
		helpView,
	)

	return view + "\n\n" + m.commandline.View()
}

func main() {
	tw, err := tw.NewTaskWarrior()
	if err != nil {
		panic(err)
	}

	tw.AddFilter("status", "pending")

	columns := []string{"ID", "Project", "Tags", "Description", "Status"}

	m := InitModel(tw, columns)
	if _, err := tea.NewProgram(m, tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

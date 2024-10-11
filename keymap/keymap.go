package keymap

import "github.com/charmbracelet/bubbles/key"

// keyMap defines a set of keybindings. To work for help it must satisfy
// key.Map. It could also very easily be a map[string]key.Binding.
type keyMap struct {
	Up      key.Binding
	Down    key.Binding
	Left    key.Binding
	Right   key.Binding
	Command key.Binding
	Space   key.Binding
	Done    key.Binding
	Undo    key.Binding
	Next    key.Binding
	Delete  key.Binding
	Search  key.Binding
	Filter  key.Binding
	Help    key.Binding
	Quit    key.Binding
}

// ShortHelp returns keybindings to be shown in the mini help view. It's part
// of the key.Map interface.
func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Command, k.Help, k.Quit}
}

// FullHelp returns keybindings for the expanded help view. It's part of the
// key.Map interface.
func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Space, k.Done, k.Undo, k.Next},
		{k.Filter, k.Left, k.Right, k.Delete},
		{k.Up, k.Down},
		{k.Help, k.Quit},
	}
}

var KeyMap = keyMap{
	Up: key.NewBinding(
		key.WithKeys("up", "k"),
		key.WithHelp("↑/k", "move up"),
	),
	Down: key.NewBinding(
		key.WithKeys("down", "j"),
		key.WithHelp("↓/j", "move down"),
	),
	Left: key.NewBinding(
		key.WithKeys("left", "h"),
		key.WithHelp("←/h", "select filter left"),
	),
	Right: key.NewBinding(
		key.WithKeys("right", "l"),
		key.WithHelp("→/l", "select filter right"),
	),
	Space: key.NewBinding(
		key.WithKeys(" "),
		key.WithHelp("space", "start/stop task"),
	),
	Command: key.NewBinding(
		key.WithKeys(":"),
		key.WithHelp(":", "enter command"),
	),
	Done: key.NewBinding(
		key.WithKeys("d"),
		key.WithHelp("d", "mark task as done"),
	),
	Undo: key.NewBinding(
		key.WithKeys("u"),
		key.WithHelp("u", "undo last action"),
	),
	Next: key.NewBinding(
		key.WithKeys("n"),
		key.WithHelp("n", "add +next tag to task"),
	),
	Delete: key.NewBinding(
		key.WithKeys("d"),
		key.WithHelp("d", "delete filter"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "toggle help"),
	),
	Search: key.NewBinding(
		key.WithKeys("/"),
		key.WithHelp("/", "search tasks"),
	),
	Filter: key.NewBinding(
		key.WithKeys("f"),
		key.WithHelp("f", "filter tasks"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "esc", "ctrl+c"),
		key.WithHelp("esc/q", "quit"),
	),
}

# taskwarrior-tui

Another Taskwarrior TUI written in Go with Bubble Tea.
It's heavily inspired by [taskwarrior-tui](https://github.com/kdheepak/taskwarrior-tui) by [kdheepak](https://github.com/kdheepak), but more lightweight and with a different goal.

## Features
- vim-like navigation
- Display tasks
- Mark tasks as done (status: completed)
- Mark tasks as next
- Start/Stop tasks
- Undo the last action
- do all the other taskwarrior commands with auto-completion

### Planned Features
- Context switching
- Project switching
- Tag management
- Task editing

![](https://github.com/Piitschy/taskwarrior-tui/blob/main/docs/assets/taskwarrior-tui.gif)

## Installation

It is necessary to have taskwarrior installed on your system. If not, you can download it [here](https://taskwarrior.org/download/).

### Go
Make sure you have Go installed on your system. If not, you can download it [here](https://golang.org/dl/).

```sh 
go install github.com/Piitschy/taskwarrior-tui@latest
```

### Executable for Linux, MacOS and Windows
Or just use the executable from the [releases](https://github.com/Piitschy/taskwarrior-tui/releases).

## Usage

```sh 
taskwarrior-tui
```

## Keybindings

### Task Table
- `:` to open the command prompt that starts with `task ` to execute taskwarrior commands
- `a` to add a new task
- `j`/`k` or `↓`/`↑` to navigate the tasks
- `q` to quit
- `d` mark task as done
- `n` to mark task as next
- `space` to start/stop task
- `u` to undo the last action
- `f` to add a new filter 
- `h`/`l` or `←`/`→` to switch to the filter section

### Filter Section
- `h`/`l` or `←`/`→` to navigate the filter
- `d` to delete filter
- `x` to disable filter
- `q`, `j`/`k` or `↓`/`↑` to close the filter section



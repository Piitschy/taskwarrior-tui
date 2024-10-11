# taskworrior-tui

Another Taskwarrior TUI written in Go with Bubble Tea.
It's heavily inspired by [taskwarrior-tui](https://github.com/kdheepak/taskwarrior-tui) by [kdheepak](https://github.com/kdheepak), but more lightweight and with a different goal.

## Features
- Display tasks
- Mark tasks as done
- Mark tasks as next
- Start/Stop tasks
- Undo the last action
- do all the other taskwarrior commands with auto-completion

### Planned Features
- Context switching
- Project switching
- Tag management
- Task editing

## Installation

```sh 
go install github.com/Piitschy/taskwarrior-tui@latest
```

## Usage

```sh 
taskwarrior-tui
```

## Keybindings

### Task Table
- `:` to open the command prompt that starts with `task ` to execute taskwarrior commands
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
- `q`, `j`/`k` or `↓`/`↑` to close the filter section



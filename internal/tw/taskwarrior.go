package tw

import (
	"encoding/json"
	"os/exec"
	"slices"
	"strconv"
)

type TaskWarrior struct {
	Tasks
	Context        string
	filter         Filters
	OnFilterChange func()
}

func NewTaskWarrior() (*TaskWarrior, error) {
	tw := TaskWarrior{}
	err := tw.LoadTasks()
	if err != nil {
		return nil, err
	}
	return &tw, nil
}

func (tw *TaskWarrior) GetTasks() Tasks {
	tasks, err := tw.GetNextTasks()
	if err != nil {
		return tw.Tasks
	}
	for _, task := range tw.Tasks {
		if !tasks.Contains(task.Id) {
			tasks = append(tasks, task)
		}
	}
	return tasks
}

func (tw *TaskWarrior) LoadTasks() error {
	tasksString, err := exec.Command("task", "export").Output()
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(tasksString), &tw.Tasks)
	return nil
}

func (tw *TaskWarrior) GetFilters() Filters {
	return tw.filter
}

func (tw *TaskWarrior) SetFilter(f Filters) {
	tw.filter = f
	tw.OnFilterChange()
}

func (tw *TaskWarrior) AddFilter(key, value string) {
	tw.filter.AddFilter(key, value)
	tw.OnFilterChange()
}

func (tw *TaskWarrior) AddFilterFromString(filterString string) error {
	err := tw.filter.AddFilterFromString(filterString)
	tw.OnFilterChange()
	return err
}

func (tw *TaskWarrior) ToggleDisableFilter(f Filter) {
	for i, filter := range tw.filter {
		if filter.key == f.key && filter.value == f.value {
			tw.filter[i].Disabled = !tw.filter[i].Disabled
		}
	}
	tw.OnFilterChange()
}

func (tw *TaskWarrior) RemoveFilter(f Filter) {
	var newFilters Filters
	for i, filter := range tw.filter {
		if filter.key != f.key || filter.value != f.value {
			newFilters = append(newFilters, tw.filter[i])
		}
	}
	tw.filter = newFilters
	tw.OnFilterChange()
}

func (tw *TaskWarrior) GetProjects() []string {
	projects := []string{}
	for _, task := range tw.Tasks {
		if task.Project != "" && !slices.Contains(projects, task.Project) {
			projects = append(projects, task.Project)
		}
	}
	return projects
}

func (tw *TaskWarrior) GetTaskById(id int) (*Task, error) {
	for _, task := range tw.Tasks {
		if task.Id == id {
			return &task, nil
		}
	}
	return nil, nil
}

func (tw *TaskWarrior) GetFilteredTasks() Tasks {
	tasks := tw.GetTasks()
	if len(tasks) == 0 {
		return tasks
	}
	for _, filter := range tw.filter {
		if filter.Disabled {
			continue
		}
		tasks = tasks.Filter(filter)
	}
	return tasks
}

func (tw *TaskWarrior) GetActiveTasks() (Tasks, error) {
	tasks := Tasks{}
	tasksString, err := exec.Command("task", "export", "active").Output()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(tasksString), &tasks)
	return tasks, nil
}

func (tw *TaskWarrior) GetNextTasks() (Tasks, error) {
	tasks := Tasks{}
	tasksString, err := exec.Command("task", "+next", "export").Output()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(tasksString), &tasks)
	return tasks, nil
}

func (tw *TaskWarrior) EditTask(taskId int) error {
	_, err := exec.Command("task", strconv.Itoa(taskId), "edit").Output()
	if err != nil {
		return err
	}
	return nil
}

func (tw *TaskWarrior) TaskDone(taskId int) error {
	_, err := exec.Command("task", strconv.Itoa(taskId), "done").Output()
	if err != nil {
		return err
	}
	return nil
}

func (tw *TaskWarrior) Undo() error {
	_, err := exec.Command("task", "rc.confirmation=off", "undo").Output()
	return err
}

func (tw *TaskWarrior) StartTask(taskId int) error {
	_, err := exec.Command("task", strconv.Itoa(taskId), "start").Output()
	if err != nil {
		return err
	}
	return nil
}

func (tw *TaskWarrior) StopTask(taskId int) error {
	_, err := exec.Command("task", strconv.Itoa(taskId), "stop").Output()
	if err != nil {
		return err
	}
	return nil
}

func (tw *TaskWarrior) TaskNext(taskId int) error {
	_, err := exec.Command("task", strconv.Itoa(taskId), "modify", "+next").Output()
	if err != nil {
		return err
	}
	return nil
}

func (tw *TaskWarrior) TaskUnnext(taskId int) error {
	_, err := exec.Command("task", strconv.Itoa(taskId), "modify", "-next").Output()
	if err != nil {
		return err
	}
	return nil
}

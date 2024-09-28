package tw

import (
	"encoding/json"
	"os/exec"
	"strconv"
)

type TaskWarrior struct {
	Tasks
	Context string
}

func NewTaskWarrior() (*TaskWarrior, error) {
	tw := TaskWarrior{}
	err := tw.LoadTasks()
	if err != nil {
		return nil, err
	}
	return &tw, nil
}

func (tw *TaskWarrior) LoadTasks() error {
	tasksString, err := exec.Command("task", "export").Output()
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(tasksString), &tw.Tasks)
	return nil
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
	_, err := exec.Command("task", strconv.Itoa(taskId), "modify +next").Output()
	if err != nil {
		return err
	}
	return nil
}

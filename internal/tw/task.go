package tw

import (
	"reflect"
	"strings"

	"github.com/fatih/structs"
)

type Task struct {
	Id          int      `json:"id"`
	Description string   `json:"description"`
	Project     string   `json:"project"`
	Status      string   `json:"status"`
	Uuid        string   `json:"uuid"`
	Urgency     float32  `json:"urgency"`
	Priority    string   `json:"priority"`
	Due         string   `json:"due"`
	End         string   `json:"end"`
	Entry       string   `json:"entry"`
	Modified    string   `json:"modified"`
	Tags        []string `json:"tags"`
}

type Tasks []Task

func (t Tasks) Len() int {
	return len(t)
}

func (t *Tasks) Contains(id int) bool {
	for _, task := range *t {
		if task.Id == id {
			return true
		}
	}
	return false
}

func (t Tasks) GetIds() []int {
	ids := make([]int, len(t))
	for i, task := range t {
		ids[i] = task.Id
	}
	return ids
}

func (t Tasks) Format(fields ...string) [][]string {
	rows := make([][]string, len(t))
	r := reflect.ValueOf(t)
	for i := 0; i < r.Len(); i++ {
		row := make([]string, len(fields))
		for j, field := range fields {
			row[j] = r.Index(i).FieldByName(field).String()
		}
		rows[i] = row
	}
	return rows
}

func (t Tasks) Filter(f Filter) Tasks {
	var tasks Tasks
	for _, task := range t {
		for key, value := range structs.Map(task) {
			if strings.Contains(f.key, ".not") {
				if l(key) == l(strings.ReplaceAll(f.key, ".not", "")) && l(value.(string)) != l(f.value) {
					tasks = append(tasks, task)
				}
			} else {
				if l(key) == l(f.key) && l(value.(string)) == l(f.value) {
					tasks = append(tasks, task)
				}
			}
		}
	}
	return tasks
}

func (t Tasks) FilterByProject(project string) Tasks {
	var tasks Tasks
	for _, task := range t {
		if task.Project == project {
			tasks = append(tasks, task)
		}
	}
	return tasks
}

func (t Tasks) FilterByStatus(status string) Tasks {
	var tasks Tasks
	for _, task := range t {
		if task.Status == status {
			tasks = append(tasks, task)
		}
	}
	return tasks
}

func (t Tasks) FilterPending() Tasks {
	return t.FilterByStatus("pending")
}

func (t Tasks) FilterCompleted() Tasks {
	return t.FilterByStatus("completed")
}

func (t Tasks) FilterDeleted() Tasks {
	return t.FilterByStatus("deleted")
}

func (t Tasks) SortIdsFirst(ids ...int) Tasks {
	var tasks Tasks
	for _, id := range ids {
		for _, task := range t {
			if task.Id == id {
				tasks = append(tasks, task)
			}
		}
	}
	for _, task := range t {
		if !tasks.Contains(task.Id) {
			tasks = append(tasks, task)
		}
	}
	return tasks
}

func l(s string) string {
	return strings.ToLower(s)
}

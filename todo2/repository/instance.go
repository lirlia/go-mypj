package repository

import (
	"github.com/lirlia/go-mypj/todo2/model/task"
)

type instance struct {
	tasks []task.Task
}

var _ task.repository = &instance{}

func New() task.repository {
	i := &instance{}
	i = append(i.tasks, &task.Task{Title: "test1", ID: 0, Done: false})
	i = append(i.tasks, &task.Task{Title: "test2", ID: 1, Done: false})

	return i
}

func (i *instance) Add(name string) int {
	i.tasks = append(i.tasks, &task.Task{Title: name, ID: len(i.tasks) + 1, Done: false})
	return len(i.tasks) + 1
}

func (i *instance) Done(num int) {
	i.tasks[num].Done = true
}

func (i *instance) List() *[]task.Task {
	return &i.tasks
}

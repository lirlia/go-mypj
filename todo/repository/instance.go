package repository

import (
	"github.com/lirlia/go-mypj/todo/model/task"
)

type instance struct {∏
	tasks []task.Task
}

func New() task.Repository {
	s := new(instance)
	s.tasks = make([]task.Task, 2, 20)

	s.tasks[0] = task.Task{
		ID:    1,
		Title: "task1",
		Done:  false,
	}
	s.tasks[1] = task.Task{
		ID:    2,
		Title: "task2",
		Done:  true,
	}
	return s
}

func (s *instance) Add(task task.Task) int {
	task.ID = len(s.tasks) + 1
	s.tasks = append(s.tasks, task)
	return task.ID
}

func (s *instance) Done(id int) {
	s.tasks[id].Done = true
}

func (s *instance) List() *[]task.Task {
	return &s.tasks
}

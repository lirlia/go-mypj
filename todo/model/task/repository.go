package task

type Repository interface {
	Add(Task) int
	Done(id int)
	List() *[]Task
}

package task

type Repository interface {
	Add(name string) int
	Done(int)
	List() *[]task.Task
}

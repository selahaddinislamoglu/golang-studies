package handler

type Handler interface {
	AddTask(description string) (int, error)
	DeleteTask(id string) error
	UpdateTaskDescription(id string, description string) error
	UpdateTaskStatus(id string, status string) error
}

package handler

type Handler interface {
	AddTask(description string) (int, error)
	DeleteTask(id string) error
	UpdateTaskDescription(id string, description string) error
	UpdateTaskStatusInProgress(id string) error
	UpdateTaskStatusDone(id string) error
	UpdateTaskStatusTodo(id string) error
	ListTasks() ([][]byte, error)
	ListTasksByStatus(status string) ([][]byte, error)
}

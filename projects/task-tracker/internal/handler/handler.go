package handler

type Handler interface {
	AddTask(description string) (int, error)
}

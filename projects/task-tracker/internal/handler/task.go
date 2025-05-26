package handler

import (
	"errors"

	"github.com/selahaddinislamoglu/golang-studies/projects/task-tracker/internal/service"
)

type TaskHandler struct {
	service service.Service
}

var ErrEmptyDescription = errors.New("task description cannot be empty")

func NewTaskHandler(service service.Service) *TaskHandler {
	return &TaskHandler{
		service: service,
	}
}

func (h *TaskHandler) AddTask(description string) (int, error) {

	if description == "" {
		return 0, ErrEmptyDescription
	}

	task, err := h.service.AddTask(description)
	if err != nil {
		return 0, err
	}
	return task.ID, nil
}

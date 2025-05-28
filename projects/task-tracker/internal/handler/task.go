package handler

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/selahaddinislamoglu/golang-studies/projects/task-tracker/internal/service"
)

type TaskHandler struct {
	service service.Service
}

var ErrEmptyDescription = errors.New("task description cannot be empty")
var ErrEmptyID = errors.New("id cannot be empty")

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

func (h *TaskHandler) DeleteTask(id string) error {

	if id == "" {
		return ErrEmptyID
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return fmt.Errorf("failed to convert id to integer: %w", err)
	}

	return h.service.DeleteTaskByID(idInt)
}

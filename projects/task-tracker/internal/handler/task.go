package handler

import (
	"encoding/json"
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

func (h *TaskHandler) UpdateTaskDescription(id, description string) error {
	if id == "" {
		return ErrEmptyID
	}

	if description == "" {
		return ErrEmptyDescription
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return fmt.Errorf("failed to convert id to integer: %w", err)
	}

	return h.service.UpdateTaskDescription(idInt, description)
}

func (h *TaskHandler) updateTaskStatus(id, status string) error {
	if id == "" {
		return ErrEmptyID
	}

	if status == "" {
		return fmt.Errorf("status cannot be empty")
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return fmt.Errorf("failed to convert id to integer: %w", err)
	}

	return h.service.UpdateTaskStatus(idInt, status)
}

func (h *TaskHandler) UpdateTaskStatusInProgress(id string) error {
	return h.updateTaskStatus(id, "in-progress")
}

func (h *TaskHandler) UpdateTaskStatusDone(id string) error {
	return h.updateTaskStatus(id, "done")
}

func (h *TaskHandler) UpdateTaskStatusTodo(id string) error {
	return h.updateTaskStatus(id, "todo")
}

func (h *TaskHandler) ListTasks() ([][]byte, error) {
	tasks, err := h.service.ListTasks()
	if err != nil {
		return nil, fmt.Errorf("failed to list tasks: %w", err)
	}

	var result [][]byte
	for _, task := range tasks {
		data, err := json.MarshalIndent(task, "", "  ")
		if err != nil {
			return nil, fmt.Errorf("failed to marshal task: %w", err)
		}
		result = append(result, data)
	}
	return result, nil
}

func (h *TaskHandler) ListTasksByStatus(status string) ([][]byte, error) {

	if status == "" {
		return nil, fmt.Errorf("status cannot be empty")
	}

	tasks, err := h.service.ListTasksByStatus(status)
	if err != nil {
		return nil, fmt.Errorf("failed to list tasks by status: %w", err)
	}

	var result [][]byte
	for _, task := range tasks {
		data, err := json.MarshalIndent(task, "", "  ")
		if err != nil {
			return nil, fmt.Errorf("failed to marshal task: %w", err)
		}
		result = append(result, data)
	}
	return result, nil
}

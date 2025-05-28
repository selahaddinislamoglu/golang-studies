package service

import (
	"errors"
	"time"

	"github.com/selahaddinislamoglu/golang-studies/projects/task-tracker/internal/model"
	"github.com/selahaddinislamoglu/golang-studies/projects/task-tracker/internal/repository"
)

// Service defines the interface for task-related operations.
type TaskService struct {
	repository repository.Repository
}

var ErrTaskAlreadyExists = errors.New("task already exists")
var ErrTaskNotFound = errors.New("task not found")

func NewTaskService(repository repository.Repository) Service {
	return &TaskService{
		repository: repository,
	}
}

// CreateTask creates a new task with the given description.
func (s *TaskService) AddTask(description string) (*model.Task, error) {

	task := s.repository.GetTaskByDescription(description)
	if task != nil {
		return nil, ErrTaskAlreadyExists
	}

	task = &model.Task{
		Description: description,
		Status:      "todo",
		CreatedAt:   time.Now().Format(time.RFC3339),
		UpdatedAt:   time.Now().Format(time.RFC3339),
	}

	err := s.repository.AddTask(task)
	if err != nil {
		return nil, err
	}
	return task, nil
}

// DeleteTaskByID deletes a task by its ID.
func (s *TaskService) DeleteTaskByID(id int) error {
	task, err := s.repository.GetTaskByID(id)
	if err != nil {
		return err
	}
	if task == nil {
		return ErrTaskNotFound
	}

	return s.repository.DeleteTask(task)
}

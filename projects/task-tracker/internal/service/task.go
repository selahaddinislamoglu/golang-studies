package service

import (
	"errors"

	"github.com/selahaddinislamoglu/golang-studies/projects/task-tracker/internal/model"
	"github.com/selahaddinislamoglu/golang-studies/projects/task-tracker/internal/repository"
)

// Service defines the interface for task-related operations.
type TaskService struct {
	repository repository.Repository
}

var ErrTaskAlreadyExists = errors.New("task already exists")
var ErrTaskNotFound = errors.New("task not found")
var ErrNoTasksFound = errors.New("no tasks found")

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

func (s *TaskService) UpdateTaskDescription(id int, description string) error {

	existingTask, err := s.repository.GetTaskByID(id)
	if err != nil {
		return err
	}
	if existingTask == nil {
		return ErrTaskNotFound
	}
	existingTask.Description = description
	return s.repository.UpdateTask(existingTask)
}

func (s *TaskService) UpdateTaskStatus(id int, status string) error {

	if status != "todo" && status != "in-progress" && status != "done" {
		return errors.New("invalid status: must be 'todo', 'in-progress', or 'done'")
	}
	existingTask, err := s.repository.GetTaskByID(id)
	if err != nil {
		return err
	}
	if existingTask == nil {
		return ErrTaskNotFound
	}
	existingTask.Status = status
	return s.repository.UpdateTask(existingTask)
}

func (s *TaskService) ListTasks() ([]*model.Task, error) {
	tasks, err := s.repository.GetAllTasks()
	if err != nil {
		return nil, err
	}
	if tasks == nil {
		return nil, ErrNoTasksFound
	}
	return tasks, nil
}

func (s *TaskService) ListTasksByStatus(status string) ([]*model.Task, error) {
	if status != "todo" && status != "in-progress" && status != "done" {
		return nil, errors.New("invalid status: must be 'todo', 'in-progress', or 'done'")
	}
	tasks, err := s.repository.GetAllTasks()
	if err != nil {
		return nil, err
	}
	if tasks == nil {
		return nil, ErrNoTasksFound
	}
	var filteredTasks []*model.Task
	for _, task := range tasks {
		if task.Status == status {
			filteredTasks = append(filteredTasks, task)
		}
	}
	if len(filteredTasks) == 0 {
		return nil, ErrNoTasksFound
	}
	return filteredTasks, nil
}

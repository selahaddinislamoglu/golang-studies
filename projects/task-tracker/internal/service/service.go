package service

import (
	"github.com/selahaddinislamoglu/golang-studies/projects/task-tracker/internal/model"
)

type Service interface {
	AddTask(description string) (*model.Task, error)
	DeleteTaskByID(id int) error
	UpdateTaskDescription(id int, description string) error
	UpdateTaskStatus(id int, status string) error
	ListTasks() ([]*model.Task, error)
	ListTasksByStatus(status string) ([]*model.Task, error)
}

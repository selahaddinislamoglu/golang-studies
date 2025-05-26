package repository

import (
	"github.com/selahaddinislamoglu/golang-studies/projects/task-tracker/internal/model"
)

type Repository interface {
	AddTask(task *model.Task) error
	GetTaskByDescription(description string) *model.Task
}

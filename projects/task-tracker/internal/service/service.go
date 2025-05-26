package service

import (
	"github.com/selahaddinislamoglu/golang-studies/projects/task-tracker/internal/model"
)

type Service interface {
	AddTask(description string) (*model.Task, error)
}

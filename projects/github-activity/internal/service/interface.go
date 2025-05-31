package service

import "github.com/selahaddinislamoglu/golang-studies/projects/github-activity/internal/model"

type ServiceInterface interface {
	GetUserActivity(username string) ([]model.Event, error)
}

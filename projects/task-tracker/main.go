package main

import (
	"github.com/selahaddinislamoglu/golang-studies/projects/task-tracker/cmd"
	"github.com/selahaddinislamoglu/golang-studies/projects/task-tracker/internal/handler"
	"github.com/selahaddinislamoglu/golang-studies/projects/task-tracker/internal/repository"
	"github.com/selahaddinislamoglu/golang-studies/projects/task-tracker/internal/service"
)

func main() {

	repository, err := repository.NewFileRepository("tasks.json")
	if err != nil {
		panic(err)
	}
	service := service.NewTaskService(repository)
	handler := handler.NewTaskHandler(service)
	cmd := cmd.NewRootCmd(handler)
	cmd.Execute()
}

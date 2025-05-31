package main

import (
	"github.com/selahaddinislamoglu/golang-studies/projects/github-activity/cmd"
	"github.com/selahaddinislamoglu/golang-studies/projects/github-activity/internal/handler"
	"github.com/selahaddinislamoglu/golang-studies/projects/github-activity/internal/service"
)

func main() {

	activityService := service.NewActivityService()
	activityHandler := handler.NewActivityHandler(activityService)
	cmd := cmd.NewRootCmd(activityHandler)
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}

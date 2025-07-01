package main

import (
	"github.com/selahaddinislamoglu/golang-studies/projects/personal-blog/internal/controller"
	"github.com/selahaddinislamoglu/golang-studies/projects/personal-blog/internal/repository"
	"github.com/selahaddinislamoglu/golang-studies/projects/personal-blog/internal/router"
	"github.com/selahaddinislamoglu/golang-studies/projects/personal-blog/internal/server"
	"github.com/selahaddinislamoglu/golang-studies/projects/personal-blog/internal/service"
)

func main() {
	repository, err := repository.NewFileRepository("./articles/")
	if err != nil {
		panic(err)
	}

	service := service.NewPersonalBlogService(repository)
	httpController := controller.NewPersonalBlogHTTPController(service)
	httpRouter := router.NewPersonalBlogHTTPRouter(httpController)
	server := server.NewPersonalBlogHTTPServer(httpRouter)

	err = server.Serve(":8080")
	if err != nil {
		panic(err)
	}
	err = server.Shutdown()
	if err != nil {
		panic(err)
	}
}

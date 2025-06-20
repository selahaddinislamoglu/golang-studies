package main

import (
	"github.com/selahaddinislamoglu/golang-studies/projects/unit-converter/internal/controller"
	"github.com/selahaddinislamoglu/golang-studies/projects/unit-converter/internal/router"
	"github.com/selahaddinislamoglu/golang-studies/projects/unit-converter/internal/server"
	"github.com/selahaddinislamoglu/golang-studies/projects/unit-converter/internal/service"
)

func main() {
	service := service.NewUnitConverterService()
	controller := controller.NewUnitConverterController(service)
	router := router.NewUnitConverterRouter(controller)
	server := server.NewUnitConverterServer(router)
	err := server.Serve(":8080")
	if err != nil {
		panic(err)
	}
	server.Shutdown()
}

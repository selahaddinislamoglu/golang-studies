package router

import (
	"net/http"

	"github.com/selahaddinislamoglu/golang-studies/projects/unit-converter/internal/controller"
)

type UnitConverterRouter struct {
	controller controller.Controller
}

func NewUnitConverterRouter(controller controller.Controller) Router {
	return &UnitConverterRouter{
		controller: controller,
	}
}

func (r *UnitConverterRouter) SetupRoutes() (http.Handler, error) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", r.controller.Home)
	mux.HandleFunc("/length", r.controller.Length)
	mux.HandleFunc("/weight", r.controller.Weight)
	mux.HandleFunc("/temperature", r.controller.Temperature)

	return mux, nil
}

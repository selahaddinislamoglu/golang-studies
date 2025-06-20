package server

import (
	"net/http"

	"github.com/selahaddinislamoglu/golang-studies/projects/unit-converter/internal/router"
)

type UnitConverterServer struct {
	router router.Router
}

func NewUnitConverterServer(router router.Router) Server {
	return &UnitConverterServer{
		router: router,
	}
}

func (s *UnitConverterServer) Serve(addr string) error {
	handler, err := s.router.SetupRoutes()
	if err != nil {
		return err
	}

	return http.ListenAndServe(addr, handler)
}

func (s *UnitConverterServer) Shutdown() error {
	return nil
}

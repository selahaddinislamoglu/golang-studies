package server

import (
	"net/http"

	"github.com/selahaddinislamoglu/golang-studies/projects/personal-blog/internal/router"
)

type PersonalBlogHTTPServer struct {
	router router.HTTPRouter
}

func NewPersonalBlogHTTPServer(router router.HTTPRouter) HTTPServer {
	return &PersonalBlogHTTPServer{
		router: router,
	}
}

func (s *PersonalBlogHTTPServer) Serve(addr string) error {
	handler, err := s.router.SetupRoutes()
	if err != nil {
		return err
	}

	return http.ListenAndServe(addr, handler)
}

func (s *PersonalBlogHTTPServer) Shutdown() error {
	return nil
}

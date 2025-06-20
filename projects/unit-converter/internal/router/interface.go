package router

import "net/http"

type Router interface {
	SetupRoutes() (http.Handler, error)
}

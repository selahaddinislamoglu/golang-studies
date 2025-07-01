package router

import "net/http"

type HTTPRouter interface {
	SetupRoutes() (http.Handler, error)
}

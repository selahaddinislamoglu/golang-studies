package controller

import "net/http"

type Controller interface {
	Home(response http.ResponseWriter, request *http.Request)
	Length(response http.ResponseWriter, request *http.Request)
	Weight(response http.ResponseWriter, request *http.Request)
	Temperature(response http.ResponseWriter, request *http.Request)
}

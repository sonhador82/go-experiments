package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)


func main() {

	unauthenticatedRouter := NewRouter()
	unauthenticatedRouter.GET("/", HandleHome)
	
	authenticatedRouter := NewRouter()
	authenticatedRouter.GET("/images/new", HandleImageNew)

	middleware := Middleware{}
	middleware.Add(unauthenticatedRouter)
	middleware.Add(http.HandlerFunc(AuthenticateRequest))
	middleware.Add(authenticatedRouter)
	http.ListenAndServe(":8080", middleware)
}

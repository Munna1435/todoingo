package main

import (
	"github.com/gorilla/mux"
	"todoingo/services/todoservice"
)

func setupRouter() *mux.Router {
	r := mux.NewRouter()
	baseRouter := r.PathPrefix("/api/v1").Subrouter()

	todoBaseRouter := baseRouter.PathPrefix("/users/{userId}/todos").Subrouter()
	todoservice.SetUpTodoRouter(todoBaseRouter)

	return r
}

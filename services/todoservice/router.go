package todoservice

import (
	"github.com/gorilla/mux"
	"net/http"
	"todoingo/repositories/store/tododatastore"
)

func SetUpTodoRouter(r *mux.Router) {
	repo := tododatastore.New()
	service := New(repo)

	r.Methods(http.MethodGet).Path("").HandlerFunc(service.GetAllTodos)
	r.Methods(http.MethodGet).Path("/{todoId}").HandlerFunc(service.GetTodo)
	r.Methods(http.MethodPost).Path("").HandlerFunc(service.CreateTodo)
	r.Methods(http.MethodPut).Path("/{todoId}").HandlerFunc(service.UpdateTodo)
	r.Methods(http.MethodDelete).Path("/{todoId}").HandlerFunc(service.DeleteTodo)
}

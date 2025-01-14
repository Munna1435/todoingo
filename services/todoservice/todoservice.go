package todoservice

import (
	"github.com/gorilla/mux"
	json "github.com/json-iterator/go"
	"net/http"
	"todoingo/models"
	"todoingo/repositories/store/tododatastore"
)

type ToDoService struct {
	repo *tododatastore.TodoDataStore
}

func New(repo *tododatastore.TodoDataStore) *ToDoService {
	return &ToDoService{repo: repo}
}

func (s *ToDoService) GetAllTodos(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	todos, err := s.repo.GetAllTodos(params["userId"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func (s *ToDoService) GetTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	todo, err := s.repo.GetTodo(params["userId"], params["todoId"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&todo)
}

func (s *ToDoService) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.TodoData
	_ = json.NewDecoder(r.Body).Decode(&todo)

	todoEntity, err := s.repo.CreateToDo(todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todoEntity)
}

func (s *ToDoService) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var todo models.TodoData
	_ = json.NewDecoder(r.Body).Decode(&todo)
	todo.Id = params["todoId"]
	todo.UserId = params["userId"]
	err := s.repo.UpdateToDo(todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func (s *ToDoService) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var todo models.TodoData
	todo.Id = params["todoId"]
	todo.UserId = params["userId"]
	err := s.repo.DeleteToDo(todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

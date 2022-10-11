package api

import (
	"github.com/Tushar-987/todo/api/controllers"
	"github.com/gorilla/mux"
)

// Register all the routes for the application
var RouteTodos = func(router *mux.Router) *mux.Router {
	router.HandleFunc("/todos", controllers.GetAllTodos).Methods("GET").Name("todos.get-all")
	router.HandleFunc("/todos/{id}", controllers.GetTodoById).Methods("GET").Name("todos.get-single")
	router.HandleFunc("/todos", controllers.CreateTodo).Methods("POST").Name("todos.add-single")
	router.HandleFunc("/todos/{id}", controllers.UpdateTodo).Methods("PUT").Name("todos.update-single")
	router.HandleFunc("/todos/{id}", controllers.DeleteTodo).Methods("DELETE").Name("todos.delete-single")

	return router
}

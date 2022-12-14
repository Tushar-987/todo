package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Tushar-987/todo/api/models"
	"github.com/Tushar-987/todo/api/serializers"
	"github.com/Tushar-987/todo/utils"
)

func GetAllTodos(w http.ResponseWriter, r *http.Request) {

	todo := models.Todo{}

	// get all todos
	todos, err := todo.All()

	// if an error is found, sent the status to the client
	if err != nil {
		switch err {
		case utils.ErrResourceNotFound:
			utils.FindError(w, err, http.StatusNotFound)
		default:
			utils.FindError(w, err, http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")

	todoSerializer := serializers.TodoSerializer{
		Todos: todos,
		Many:  true,
	}

	_ = json.NewEncoder(w).Encode(todoSerializer.Serialize()["data"])
}

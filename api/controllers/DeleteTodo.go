package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Tushar-987/todo/api/models"
	"github.com/Tushar-987/todo/api/serializers"
	"github.com/Tushar-987/todo/utils"
	"github.com/gorilla/mux"
)

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	todos := models.Todo{}

	//get the slug by the parameter 'id'
	vars := mux.Vars(r)
	idString := vars["id"]
	id, _ := strconv.Atoi(idString)
	todo, err := todos.Delete(id)

	//if an error is found send it to the client and return
	if err != nil {
		switch err {
		case utils.ErrResourceNotFound:
			utils.FindError(w, err, http.StatusNotFound)
		default:
			utils.FindError(w, err, http.StatusInternalServerError)
		}
		return
	}

	//write the header to the response
	w.WriteHeader(http.StatusNoContent)

	todoSerializer := serializers.TodoSerializer{
		Todos: []*models.Todo{
			todo,
		},
		Many: false,
	}

	// send the todo to the client
	_ = json.NewEncoder(w).Encode(todoSerializer.Serialize()["data"])

}

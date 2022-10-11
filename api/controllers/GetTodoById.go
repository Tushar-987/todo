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

func GetTodoById(w http.ResponseWriter, r *http.Request) {

	todo := models.Todo{}

	//get the slug by the parameter 'id'
	vars := mux.Vars(r)
	idString := vars["id"]
	id, _ := strconv.Atoi(idString)

	todos, err := todo.Retrieve(id)

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

	//Set the Content Type of the header to application/json
	w.Header().Set("Content-Type", "application/json")

	todoSerializer := serializers.TodoSerializer{
		Todos: []*models.Todo{
			todos,
		},
		Many: false,
	}

	//Encode the created todos response to json and send it
	_ = json.NewEncoder(w).Encode(todoSerializer.Serialize()["data"])
}

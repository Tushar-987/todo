package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/Tushar-987/todo/api/models"
	"github.com/Tushar-987/todo/api/serializers"
	"github.com/Tushar-987/todo/utils"
)

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Creation Started")
	todoInstance := models.Todo{}

	// decode the request body to todo
	json.NewDecoder(r.Body).Decode(&todoInstance)

	if todoInstance.Title == "" || todoInstance.Description == "" {
		utils.FindError(w, errors.New("missing fields"), http.StatusBadRequest)
		return
	}

	todo, err := todoInstance.Insert()
	// if an error is found, send it to the client and return
	if err != nil {
		if err == utils.ErrResourceNotFound {
			utils.FindError(w, err, http.StatusNotFound)
		} else {
			utils.FindError(w, err, http.StatusInternalServerError)
		}
		return
	}
	// set header content type to application/json
	w.Header().Set("Content-Type", "application/json")

	// Todo data serialization
	todoSerializer := serializers.TodoSerializer{
		Todos: []*models.Todo{
			todo,
		},
		Many: false,
	}

	// send the created todo to the response
	_ = json.NewEncoder(w).Encode(todoSerializer.Serialize()["data"])
}

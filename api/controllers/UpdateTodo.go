package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	todo := models.Todo{}

	//get the slug by the parameter 'id'
	vars := mux.Vars(r)
	idString := vars["id"]
	id, _ := strconv.Atoi(idString)

	// get the todo with this id first
	_, err := todo.Retrieve(id)

	if err != nil {
		if err == utils.ErrResourceNotFound {
			utils.FindError(w, err, http.StatusNotFound)
		}
		return
	}

	if todo.Title == "" || todo.Description == "" {
		utils.FindError(w, errors.New("missing fields"), http.StatusBadRequest)
		return
	}

	json.NewDecoder(r.Body).Decode(&todo)

	todos, err := todo.Update(id)

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

	// //write the header to the response
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

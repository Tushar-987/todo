package serializers

import "github.com/Tushar-987/todo/api/models"

type TodoSerializer struct {
	Todos []*models.Todo
	Many  bool
}

func (serializer *TodoSerializer) Serialize() map[string]interface{} {
	serializedData := make(map[string]interface{})

	todosArray := make([]interface{}, 0)
	for _, todo := range serializer.Todos {
		todosArray = append(todosArray, map[string]interface{}{
			//"id":          todo.Id,
			"title":       todo.Title,
			"description": todo.Description,
			"completed":   todo.Completed,
		})
	}

	if serializer.Many {
		serializedData["data"] = todosArray
	} else {
		if len(todosArray) != 0 {
			serializedData["data"] = todosArray[0]
		} else {
			serializedData["data"] = make(map[string]interface{})
		}
	}

	return serializedData
}

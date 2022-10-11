package tests

import (
	"net/http"
	"testing"
)

func TestCreateTodo(t *testing.T) {
	ts := utils.TestServer{Server: Server}

	// When the request body does not contain a description
	t.Run("Required Fields not Given", func(t *testing.T) {
		reqBody := `{
			"title": "Title 1"
		}`

		statusCode, _, _ := ts.Post(t, "/todos", reqBody, "")

		if statusCode != http.StatusBadRequest {
			t.Errorf("want %d status code; got %d", http.StatusBadRequest, statusCode)
		}
	})

	// When the request body does not contain a title
	t.Run("Required Fields not Given", func(t *testing.T) {
		reqBody := `{
			"description": "Description 1"
		}`

		statusCode, _, _ := ts.Post(t, "/todos", reqBody, "")

		if statusCode != http.StatusBadRequest {
			t.Errorf("want %d status code; got %d", http.StatusBadRequest, statusCode)
		}
	})

	// When the request is valid with all fields
	t.Run("Valid request", func(t *testing.T) {
		reqBody := `{ 
				"title": "title 5",
				"description": "description 5",
				"completed" : "not yet done"
			}`

		statusCode, _, _ := ts.Post(t, "/todos", reqBody, "")

		if statusCode != http.StatusOK {
			t.Errorf("want %d status code; got %d", http.StatusOK, statusCode)
		}
	})
}

package tests

import (
	"encoding/json"
	"net/http"
	"testing"
)

func TestUpdateTodo(t *testing.T) {
	ts := utils.TestServer{Server: Server}

	t.Run("Bad request with a non-existing todo id", func(t *testing.T) {
		reqBody := `{
			"title" : "changed title",
			"description": "changed description"
		}`

		statusCode, _, _ := ts.Put(t, "/todos/90", reqBody)

		if statusCode != http.StatusNotFound {
			t.Errorf("want %d status code; got %d", http.StatusNotFound, statusCode)
		}
	})

	t.Run("Bad request without required fields", func(t *testing.T) {
		reqBody := `{
			"title": "changed title"
		}`

		statusCode, _, _ := ts.Put(t, "/todos/2", reqBody)

		if statusCode != http.StatusBadRequest {
			t.Errorf("want %d status code; got %d", http.StatusBadRequest, statusCode)
		}
	})

	t.Run("Bad request without required fields", func(t *testing.T) {
		reqBody := `{
			"description": "changed description"
		}`

		statusCode, _, _ := ts.Put(t, "/todos/2", reqBody)

		if statusCode != http.StatusBadRequest {
			t.Errorf("want %d status code; got %d", http.StatusBadRequest, statusCode)
		}
	})

	t.Run("Bad request with empty fields", func(t *testing.T) {
		reqBody := `{
			"title": "",
			"description": ""
		}`

		statusCode, _, _ := ts.Put(t, "/todos/2", reqBody)

		if statusCode != http.StatusBadRequest {
			t.Errorf("want %d status code; got %d", http.StatusBadRequest, statusCode)
		}
	})

	t.Run("Valid request", func(t *testing.T) {
		reqBody := `{
			"title": "changed title",
			"description": "changed description",
			"completed": "done"
		}`

		statusCode, _, resBody := ts.Put(t, "/todos/6", reqBody)

		if statusCode != http.StatusOK {
			t.Errorf("want %d status code; got %d", http.StatusOK, statusCode)
		} else {
			var response map[string]interface{}

			err := json.Unmarshal(resBody, &response)
			if err != nil {
				t.Fatal("Error unmarshalling response body: ", err.Error())
			}

			if response["title"] != "changed title" {
				t.Errorf("want %s as title; got %s", "changed title", response["title"])
			}

			if response["description"] != "changed description" {
				t.Errorf("want %s as description; got %s", "changed description", response["description"])
			}

			if response["completed"] != "done" {
				t.Errorf("want %s as completed; got %s", "changed completed", response["completed"])
			}
		}
	})
}

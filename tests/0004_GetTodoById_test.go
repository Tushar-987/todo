package tests

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/Tushar-987/todo/utils"
)

func TestGetTodoById(t *testing.T) {
	ts := utils.TestServer{Server: Server}

	t.Run("Bad request with a non-existing todo", func(t *testing.T) {
		statusCode, _, _ := ts.Get(t, "/todos/90", "")

		if statusCode != http.StatusNotFound {
			t.Errorf("want %d status code; got %d", http.StatusNotFound, statusCode)
		}
	})

	t.Run("Valid request", func(t *testing.T) {
		statusCode, _, resBody := ts.Get(t, "/todos/1", "")

		if statusCode != http.StatusOK {
			t.Errorf("want %d status code; got %d", http.StatusOK, statusCode)
		} else {
			var response map[string]interface{}

			err := json.Unmarshal(resBody, &response)
			if err != nil {
				t.Fatal("Error unmarshalling response body: ", err.Error())
			}

			if response["title"] != "title1" {
				t.Errorf("want %s as the title; got %s", "title1", response["title"])
			}

			if response["description"] != "description1" {
				t.Errorf("want %s as the description; got %s", "description1", response["description"])
			}
		}
	})
}

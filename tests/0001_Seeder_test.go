package tests

import (
	"testing"

	"github.com/Tushar-987/todo/cmd"
	"github.com/Tushar-987/todo/utils"
)

func TestSeeder(t *testing.T) {
	t.Run("Seed todos", func(t *testing.T) {
		err := utils.SeedTodos(cmd.DbConnection)

		if err != nil {
			t.Fatal(err)
		}
	})
}

package tests

import "testing"

func TestSeeder(t *testing.T) {
	t.Run("Seed todos", func(t *testing.T) {
		err := utils.SeedTodos(cmd.DbConnection)

		if err != nil {
			t.Fatal(err)
		}
	})
}

package models

import (
	"database/sql"

	"github.com/Tushar-987/todo/cmd"
	"github.com/Tushar-987/todo/utils"
	_ "github.com/go-sql-driver/mysql"
)

//query := `
//CREATE TABLE IF NOT EXISTS todo (
//	id INT AUTO_INCREMENT,
//	title VARCHAR(100),
//	description VARCHAR(200),
//	completed VARCHAR(100),
//	PRIMARY KEY (id)
//);
//`

// Todo struct which describes the todo table in the database
type Todo struct {
	Id          int    `db:"id" json:"id"`
	Title       string `db:"title" json:"title"`
	Description string `db:"description" json:"description"`
	Completed   string `db:"completed" json:"completed"`
}

var TotalTodosCount int

// All - Retrieves all the records from the todo Table
// Params - todos ([]*Todo)
// Returns an error, if any

func (todo *Todo) All() ([]*Todo, error) {
	todos := make([]*Todo, 0)
	/// execute the select query
	err := cmd.DbConnection.Select(&todos, "SELECT * FROM todo")

	// if an error is found, return it
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, utils.ErrResourceNotFound
		default:
			return nil, err
		}
	}
	return todos, nil
}

// Retrieve - Retrieves a record from the Todo table
// Params - todo (*Todo), id int
// returns an error, if any
func (todo *Todo) Retrieve(id int) (*Todo, error) {
	//execute the query
	err := cmd.DbConnection.Get(todo, `SELECT * FROM todo WHERE id=? LIMIT 1`, id)

	//if an error is found, return it
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return todo, utils.ErrResourceNotFound
		default:
			return todo, err
		}
	}
	return todo, nil
}

// Insert - Inserts the value in the todo table
// params - todo (*Todo)
// Returns todo and error if any
func (todo *Todo) Insert() (*Todo, error) {

	insertQuery := "insert into todo(title,description,completed) values (:title,:description,:completed)"
	// execute the insert query
	_, err := cmd.DbConnection.NamedExec(insertQuery, &todo)
	// if an error is found, return it
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, utils.ErrResourceNotFound
		}
		return nil, err
	}

	return todo, nil
}

// Update - Modifies the values of the specified record in the todo Table
// params - todo (*Todo), id int
// Returns an error if any
func (todo *Todo) Update(id int) (*Todo, error) {
	//execute the query
	_, err := cmd.DbConnection.NamedExec("UPDATE todo SET title=:title , description=:description , completed=:completed where id=:id", &todo)

	//if an error is found, return it
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, utils.ErrResourceNotFound
		default:
			return nil, err
		}
	}
	return todo, nil
}

// Delete - Deletes the specified record from the todo Table
// params - todo (*Todo), id int
// Returns an error if any
func (todo *Todo) Delete(id int) (*Todo, error) {

	err := cmd.DbConnection.Get(todo, `SELECT * FROM todo WHERE id=? LIMIT 1`, id)

	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, utils.ErrResourceNotFound
		default:
			return nil, err
		}
	}

	//execute the query
	_, err = cmd.DbConnection.Exec(`DELETE FROM todo WHERE id=?`, id)

	//if an error is found, return it
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, utils.ErrResourceNotFound
		default:
			return nil, err
		}
	}
	return todo, nil
}

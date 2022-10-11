package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
)

var DbConnection *sqlx.DB

func Connect() (*sqlx.DB, error) {
	dbUrl := fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	//fmt.Println(dbUrl)
	dbConn, err := sqlx.Connect("mysql", dbUrl)

	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	if err = dbConn.Ping(); err != nil {
		log.Fatalln(err)
		return nil, err
	}
	query := `
	CREATE TABLE IF NOT EXISTS todo (
		id INT AUTO_INCREMENT,
		title TEXT,
		description TEXT,
		completed TEXT,
		PRIMARY KEY (id)
	);
	`
	_, err = dbConn.Exec(query)
	if err != nil {
		log.Fatalln(err)
	}

	return dbConn, nil
}

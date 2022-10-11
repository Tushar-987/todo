package main

import (
	"log"
	"net/http"

	"github.com/Tushar-987/todo/api"
	"github.com/Tushar-987/todo/cmd"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	router := mux.NewRouter()

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("No .env file found")
	}

	dbConn, err := cmd.Connect()
	if err != nil {
		log.Fatalln("Error connecting db:", err.Error())
		return
	}
	cmd.DbConnection = dbConn

	api.RouteTodos(router)

	// http.Handle("/todos",)
	log.Println("Starting the server")
	log.Fatal(http.ListenAndServe(":3000", router))

}

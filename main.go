package main

import (
	"github.com/marc06210/gc-back-app/internal/db"
	"github.com/marc06210/gc-back-app/internal/todo"
	"github.com/marc06210/gc-back-app/internal/transport"
	"log"
	"os"
	"strconv"
)

func main() {
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPortString := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	if dbUsername == "" || dbPassword == "" || dbHost == "" || dbPortString == "" || dbName == "" {
		log.Fatal("database environment variables not set")
	}
	dbPort, err := strconv.Atoi(dbPortString)
	if err != nil {
		log.Fatal("database port is not a number")
	}

	database, err := db.New(dbUsername, dbPassword, dbName, dbHost, dbPort)
	defer database.Close()
	svc := todo.NewService(database)
	server := transport.NewServer(svc)

	if err != nil {
		log.Fatal(err)
	}
	_,_ = database.GetAllPublications()

	if err := server.Serve(); err != nil {
		log.Fatal(err)
	}
}

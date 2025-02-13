package main

import (
	"github.com/marc06210/gc-back-app/internal/db"
	"github.com/marc06210/gc-back-app/internal/publication"
	"github.com/marc06210/gc-back-app/internal/transport"
	"go.uber.org/zap"
	"os"
	"strconv"
)

func main() {
	logger := zap.Must(zap.NewProduction())
	if os.Getenv("APP_ENV") == "development" {
		logger = zap.Must(zap.NewDevelopment())
	}

	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPortString := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	if dbUsername == "" || dbPassword == "" || dbHost == "" || dbPortString == "" || dbName == "" {
		logger.Fatal("database environment variables not set")
	}
	dbPort, err := strconv.Atoi(dbPortString)
	if err != nil {
		logger.Fatal("database port is not a number")
	}

	database, err := db.New(dbUsername, dbPassword, dbName, dbHost, dbPort)
	defer database.Close()

	if err != nil {
		logger.Error("failed to start the database", zap.Error(err))
	}

	svc := publication.NewService(database)
	server := transport.NewServer(svc, logger)

	if err := server.Serve(); err != nil {
		logger.Error("failed to start the server", zap.Error(err))
	}
}

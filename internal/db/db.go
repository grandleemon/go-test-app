package db

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/lpernett/godotenv"
	"log"
	"os"
)

var dbConn *pgx.Conn

func Connect() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	connString := os.Getenv("DATABASE_URL")

	if connString == "" {
		log.Fatal("DATABASE_URL not set in environment")
	}

	dbConn, err = pgx.Connect(context.Background(), connString)

	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	log.Println("Connected to database")
}

func Close() {
	if dbConn != nil {
		dbConn.Close(context.Background())
	}
}

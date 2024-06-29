package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/mwdev22/ecom/app/api"
	"github.com/mwdev22/ecom/app/db"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")

	connStr := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=5432 sslmode=disable",
		dbUser, dbName, dbPass, dbHost)

	db, err := db.NewPGConnection(connStr)
	if err != nil {
		log.Fatal("db connection failed: %v", err)
	}

	fmt.Println("connected with %v", dbName)

	server := api.NewServer(":8080", nil)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

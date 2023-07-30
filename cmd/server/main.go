package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil && os.Getenv("POSTGRES_DATABASE_CONNECTION") == "" {
		log.Fatal("Error loading .env file")
	}

	log.Println("hello world!")
}

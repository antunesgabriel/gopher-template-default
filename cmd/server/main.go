package main

import (
	"database/sql"
	"github.com/antunesgabriel/gopher-template-default/internal/config"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	if os.Getenv("APP_ENV") != "development" {
		err := godotenv.Load()

		if err != nil {
			log.Fatalln("Error loading .env file: ", err.Error())
		}
	}

	if os.Getenv("DATABASE_URL") == "" {
		log.Fatalln("DATABASE_URL is required")
	}

	db, err := config.NewDB()

	if err != nil {
		log.Fatal("Error on connect db:", err.Error())
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Println(err)
		}
	}(db)

	signKey := config.SignKey(os.Getenv("JWT_SIGN_KEY"))

	server := InitServer(db, signKey)

	panic(server.Load().Run(os.Getenv("PORT")))
}

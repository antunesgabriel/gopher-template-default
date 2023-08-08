package main

import (
	"database/sql"
	"github.com/antunesgabriel/gopher-template-default/internal/config"
	"log"
)

func main() {
	env, err := config.NewEnv("")

	if err != nil {
		log.Fatalln(err)
	}

	db, err := config.NewDB(env)

	if err != nil {
		log.Fatal("Error on connect db:", err.Error())
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Println(err)
		}
	}(db)

	server := InitServer(db, env)

	panic(server.Load().Run())
}

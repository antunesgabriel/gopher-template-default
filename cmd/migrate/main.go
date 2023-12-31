package main

import (
	"flag"
	"fmt"
	"github.com/antunesgabriel/gopher-template-default/internal/config"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	_ "github.com/lib/pq" // Driver do PostgreSQL (ou o driver do seu banco de dados)
	"github.com/pressly/goose/v3"
)

func migrationDir() string {
	return filepath.Join("scripts", "sql", "migrations")
}

func down(env *config.Env) error {
	connectionURL := env.DatabaseURL

	dir := migrationDir()

	if connectionURL == "" {
		log.Fatalln("DATABASE_URL is required")
	}

	db, err := goose.OpenDBWithDriver("postgres", connectionURL)

	if err != nil {
		return err
	}

	defer db.Close()

	err = goose.Down(db, dir)

	if err != nil {
		return err
	}

	fmt.Println("Migrations applied successfully!")
	return nil
}

func up(env *config.Env) error {
	connectionURL := env.DatabaseURL

	dir := migrationDir()

	if connectionURL == "" {
		log.Fatalln("DATABASE_URL is required")
	}

	db, err := goose.OpenDBWithDriver("postgres", connectionURL)

	if err != nil {
		return err
	}

	defer db.Close()

	err = goose.Up(db, dir)
	if err != nil {
		return err
	}

	fmt.Println("Migrations applied successfully!")
	return nil
}

func create(env *config.Env, migrationName string) error {
	// Generate a timestamp to use in the filename
	timestamp := time.Now().UTC().Format("20060102150405")
	filename := fmt.Sprintf(
		"%s_%s.sql",
		timestamp,
		strings.ToLower(strings.ReplaceAll(migrationName, " ", "_")),
	)

	filePath := filepath.Join(migrationDir(), filename)

	file, err := os.Create(filePath)

	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("-- +goose Up\n\n-- +goose Down\n\n"))

	if err != nil {
		return err
	}

	fmt.Printf("Migration file %s created successfully!\n", filePath)

	return nil
}

func main() {
	var action string
	var migrationName string

	flag.StringVar(&action, "action", "up", "Actions: --action=up // create or down")
	flag.StringVar(
		&migrationName,
		"name",
		"",
		"Migration name. Exe: --action=create --name=create_users_table",
	)
	flag.Parse()

	env, err := config.NewEnv("")

	if err != nil {
		log.Fatalln(err)
	}

	if action == "create" {
		if migrationName == "" {
			log.Fatalln("--name flag is required")
		}

		create(env, migrationName)

		return
	}

	if action == "down" {
		if err := down(env); err != nil {
			log.Fatalln(err)
		}

		return
	}

	if err := up(env); err != nil {
		log.Fatalln(err)
	}
}

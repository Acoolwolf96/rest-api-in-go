package main

import (
	"database/sql"
	"log"
	"rest-api-in-go/docs"
	"rest-api-in-go/internal/database"
	"rest-api-in-go/internal/env"
	_ "github.com/mattn/go-sqlite3"

	_ "github.com/joho/godotenv/autoload"
)


// @title REST API in Go
// @version 1.0
// @description This is a sample REST API server written in Go using the Gin framework.
// @SecurityDefinitions.apikey BearerAuth
// @In header
// @Name Authorization
// @description Type "Bearer" followed by a space and then your JWT token.


type application struct {
	port int
	jwtSecret string
	models database.Models
}

func main() {
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	models := database.NewModels(db)

	app := &application{
		port: env.GetEnvInt("PORT", 4000),
		jwtSecret: env.GetEnvKey("JWT_SECRET", "defaultsecret"),
		models: models,
	}

	if err := app.serve(); err != nil {
		log.Fatal(err)
	}
}
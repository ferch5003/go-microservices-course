package main

import (
	"database/sql"
	"log"
	"net/http"

	"authentication/internal/data"
)

const _webPort = "80"

type Config struct {
	DB     *sql.DB
	Models data.Models
}

func main() {
	log.Println("Starting authentication service...")

	// TODO connect to DB.

	// Setup config.
	app := Config{}

	// Web server
	srv := &http.Server{
		Addr:    ":" + _webPort,
		Handler: app.routes(),
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Panicln(err)
	}
}

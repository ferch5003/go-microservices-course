package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"authentication/internal/data"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const _webPort = "80"

var counts int64

type Config struct {
	Repo data.Repository
}

func main() {
	log.Println("Starting authentication service...")

	// Connect to DB.
	conn := connectToDB()
	if conn == nil {
		log.Panicln("Can't connect to Postgres!")
	}

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

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func connectToDB() *sql.DB {
	dsn := os.Getenv("DSN")

	for {
		connection, err := openDB(dsn)
		if err != nil {
			log.Println("Postgres not yet ready...")
			counts++
		} else {
			log.Println("Connected to Postgres!")
			return connection
		}

		if counts > 10 {
			log.Println(err)
			return nil
		}

		log.Println("Backing off for two seconds...")
		time.Sleep(2 * time.Second)
		continue
	}
}

func (app *Config) setRepo(conn *sql.DB) {
	db := data.NewPostgresRepository(conn)

	app.Repo = db
}

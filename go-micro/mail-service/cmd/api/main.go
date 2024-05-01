package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
)

const _webPort = "80"

type Config struct {
	Mailer Mail
}

func main() {
	app := Config{
		Mailer: createMail(),
	}

	log.Println("Starting Mail Service on port", _webPort)

	srv := &http.Server{
		Addr:    ":" + _webPort,
		Handler: app.routes(),
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Panicln(err)
	}
}

func createMail() Mail {
	port, err := strconv.Atoi(os.Getenv("MAIL_PORT"))
	if err != nil {
		log.Println(err)
	}

	m := Mail{
		Domain:      os.Getenv("MAIL_DOMAIN"),
		Host:        os.Getenv("MAIL_HOST"),
		Port:        port,
		Username:    os.Getenv("MAIL_USERNAME"),
		Password:    os.Getenv("MAIL_PASSWORD"),
		Encryption:  os.Getenv("MAIL_ENCRYPTION"),
		FromAddress: os.Getenv("MAIL_ADDRESS"),
		FromName:    os.Getenv("MAIL_NAME"),
	}

	return m
}

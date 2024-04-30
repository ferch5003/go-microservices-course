package main

import (
	"log"
	"net/http"
)

const _webPort = "80"

type Config struct {
}

func main() {
	app := Config{}

	log.Println("Starting Mail Service on port", _webPort)

	srv := &http.Server{
		Addr:    ":" + _webPort,
		Handler: app.routes(),
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Panicln(err)
	}
}

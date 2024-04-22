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

	log.Println("Starting broker service on port", _webPort)

	// Define http server.
	srv := &http.Server{
		Addr:    ":" + _webPort,
		Handler: app.routes(),
	}

	// Start the server.
	if err := srv.ListenAndServe(); err != nil {
		log.Panicln(err)
	}
}

package main

import (
	"net/http"
)

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	// Implement the broker handler.
	payload := jsonResponse{
		Message: "Hit the Broker",
	}

	_ = app.writeJSON(w, http.StatusOK, payload)
}

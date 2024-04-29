package main

import (
	"net/http"

	"log-service/internal/data"
)

type JSONPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func (app *Config) WriteLog(w http.ResponseWriter, r *http.Request) {
	// read json into var
	var requestPayload JSONPayload

	_ = app.readJSON(w, r, &requestPayload)

	// insert data
	event := data.LogEntry{
		Name: requestPayload.Name,
		Data: requestPayload.Data,
	}

	if err := app.Models.LogEntry.Insert(event); err != nil {
		app.errorJSON(w, err)
		return
	}

	response := jsonResponse{
		Message: "logged",
	}

	app.writeJSON(w, http.StatusAccepted, response)
}

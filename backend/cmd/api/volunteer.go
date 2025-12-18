package main

import (
	"github.com/cconner57/adoption-os/backend/internal/data"
	"net/http"
)

func (app *application) submitVolunteerApplication(w http.ResponseWriter, r *http.Request) {
	var input data.VolunteerApplication

	// Read JSON from frontend
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	// Insert into DB
	err = app.models.Volunteers.Insert(&input)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	// Send back the created object
	err = app.writeJSON(w, http.StatusCreated, envelope{"application": input}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

package main

import (
	"net/http"
)

// Fixes app.GetSpotlightPets
func (app *application) getSpotlightPet(w http.ResponseWriter, r *http.Request) {
	// Call the DB model we made earlier
	pet, err := app.models.Pets.GetSpotlight()
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"pet": pet}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

// Fixes app.GetAllPets
// (Note: You'll need to add a GetAll() method to internal/data/pets.go later if you haven't yet.
// For now, let's just make it return a placeholder so the error goes away)
func (app *application) getAllPets(w http.ResponseWriter, r *http.Request) {
	// placeholder for now
	data := envelope{"message": "GetAllPets coming soon"}
	app.writeJSON(w, http.StatusOK, data, nil)
}

// Fixes app.GetAvailablePets
func (app *application) getAvailablePets(w http.ResponseWriter, r *http.Request) {
	// placeholder for now
	data := envelope{"message": "GetAvailablePets coming soon"}
	app.writeJSON(w, http.StatusOK, data, nil)
}

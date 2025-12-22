package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (app *application) getSpotlightPets(w http.ResponseWriter, r *http.Request) {
	pet, err := app.models.Pets.GetSpotlight()
	if err != nil {
		fmt.Println("Handler GetSpotlight Error:", err)
		app.serverErrorResponse(w, r, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(pet); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) getAllPets(w http.ResponseWriter, r *http.Request) {
	data := envelope{"message": "GetAllPets coming soon"}
	app.writeJSON(w, http.StatusOK, data, nil)
}

func (app *application) getAvailablePets(w http.ResponseWriter, r *http.Request) {
	data := envelope{"message": "GetAvailablePets coming soon"}
	app.writeJSON(w, http.StatusOK, data, nil)
}

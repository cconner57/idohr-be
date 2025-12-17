package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"idohr-be/internal/models"
)

// Application struct holds dependencies (like the DB)
// so we don't have to pass them around manually everywhere.
type Application struct {
	DB *sql.DB
}

func (app *Application) Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Pet Adoption API is Running! 🐶"))
}

func (app *Application) GetAllPets(w http.ResponseWriter, r *http.Request) {
	rows, err := app.DB.Query("SELECT id, name, species, breed, status FROM pets LIMIT 50")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rows.Close()

	var pets []models.Pet
	for rows.Next() {
		var p models.Pet
		if err := rows.Scan(&p.ID, &p.Name, &p.Species, &p.Breed, &p.Status); err != nil {
			log.Println("Error scanning row:", err)
			continue
		}
		pets = append(pets, p)
	}

	if pets == nil {
		pets = []models.Pet{}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pets)
}

func (app *Application) GetAvailablePets(w http.ResponseWriter, r *http.Request) {
	query := `
		SELECT id, name, species, breed, status 
		FROM pets 
		WHERE status = 'available' 
		LIMIT 50`

	rows, err := app.DB.Query(query)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rows.Close()

	var pets []models.Pet
	for rows.Next() {
		var p models.Pet
		if err := rows.Scan(&p.ID, &p.Name, &p.Species, &p.Breed, &p.Status); err != nil {
			log.Println("Error scanning row:", err)
			continue
		}
		pets = append(pets, p)
	}

	if pets == nil {
		pets = []models.Pet{}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pets)
}

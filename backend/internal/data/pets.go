package data // <--- Must be 'data'

import (
	"database/sql"
	"errors"
	"time"
)

// Define the Pet struct (matching your DB table)
type Pet struct {
	ID          int64     `json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	Name        string    `json:"name"`
	Species     string    `json:"species"`
	Breed       string    `json:"breed"`
	Age         int       `json:"age"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	ImageURL    string    `json:"imageUrl"` // Optional: if you added this column
}

// Define the Model wrapper
type PetModel struct {
	DB *sql.DB
}

// Add a method to get the Spotlight pet
func (m PetModel) GetSpotlight() (*Pet, error) {
	query := `
		SELECT id, created_at, name, species, breed, age, description, status 
		FROM pets 
		ORDER BY RANDOM() 
		LIMIT 1`

	var pet Pet

	// Execute query
	err := m.DB.QueryRow(query).Scan(
		&pet.ID,
		&pet.CreatedAt,
		&pet.Name,
		&pet.Species,
		&pet.Breed,
		&pet.Age,
		&pet.Description,
		&pet.Status,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("no pets found")
		}
		return nil, err
	}

	return &pet, nil
}

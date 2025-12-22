package data // <--- Must be 'data'

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

// Define the Pet struct (matching your DB table)
type SpotlightPet struct {
	ID           string          `json:"id"`
	Name         string          `json:"name"`
	Descriptions json.RawMessage `json:"descriptions"`
	Photos       json.RawMessage `json:"photos"`
}

// Define the Model wrapper
type PetModel struct {
	DB *sql.DB
}

func (m PetModel) GetSpotlight() ([]*SpotlightPet, error) {
	// This query looks for the boolean column we updated
	query := `
        SELECT id, name, COALESCE(descriptions, '{}'), COALESCE(photos, '[]')
        FROM pets
        WHERE profile_settings->>'is_spotlight_featured' = 'true'
           OR profile_settings->>'isSpotlightFeatured' = 'true'
        LIMIT 4`

	rows, err := m.DB.Query(query)
	if err != nil {
		fmt.Println("Query Error:", err)
		return nil, err
	}
	defer rows.Close()

	pets := []*SpotlightPet{}

	for rows.Next() {
		var pet SpotlightPet
		err := rows.Scan(
			&pet.ID,
			&pet.Name,
			&pet.Descriptions,
			&pet.Photos,
		)
		if err != nil {
			fmt.Println("Scan Error:", err)
			return nil, err
		}
		pets = append(pets, &pet)
	}

	return pets, nil
}

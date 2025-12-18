package data

import "database/sql"

type Models struct {
	Volunteers VolunteerModel
	Pets       PetModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Volunteers: VolunteerModel{DB: db},
		Pets:       PetModel{DB: db},
	}
}

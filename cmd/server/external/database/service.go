package database

import (
	"database/sql"
	"dentistas/internal/dentist"

	"fmt"
)

type SqlStore struct {
	*sql.DB
}

func NewDatabase(db *sql.DB) *SqlStore {
	return &SqlStore{db}
}

func (s *SqlStore) GetByID(id int) (dentist.Dentist, error) {

	var d dentist.Dentist

	query := fmt.Sprintf("SELECT * FROM dentist WHERE id = %d;", id)
	row := s.DB.QueryRow(query)
	err := row.Scan(&d.ID, &d.Apellido, &d.Nombre, &d.Matricula)

	if err != nil {
		return dentist.Dentist{}, err
	}

	return d, nil

}

func (s *SqlStore) Add(dentista dentist.Dentist) (dentist.Dentist, error) {
	//TODO implement me
	panic("implement me Add")
}

func (s *SqlStore) Update(id int, dentista dentist.Dentist) (dentist.Dentist, error) {
	//TODO implement me
	panic("implement me Update")
}

func (s *SqlStore) Delete(id int) error {
	//TODO implement me
	panic("implement me Delete")
}

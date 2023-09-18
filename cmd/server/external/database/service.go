package database

import (
	"database/sql"
	"dentistas/internal/dentist"
	"dentistas/internal/utils"
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

	query := fmt.Sprintf("SELECT * FROM dentista WHERE id = %d;", id)
	row := s.DB.QueryRow(query)
	err := row.Scan(&d.ID, &d.Apellido, &d.Nombre, &d.Matricula)

	if err != nil {
		return dentist.Dentist{}, err
	}

	return d, nil

}

func (s *SqlStore) Create(dentist dentist.Dentist) (dentist.Dentist, error) {

	query := fmt.Sprintf("INSERT INTO dentista (apellido, nombre, matricula) VALUES ('%s', '%s', '%s');",
		dentist.Apellido, dentist.Nombre, dentist.Matricula)

	result, err := s.DB.Exec(query)
	if err != nil {
		return dentist, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return dentist, err
	}

	dentist.ID = int(id)
	return dentist, nil
}

func (s *SqlStore) Update(id int, dentist dentist.Dentist) (dentist.Dentist, error) {

	query := fmt.Sprintf("UPDATE dentista SET apellido = '%s', nombre = '%s', matricula = '%s' WHERE id = %d;",
		dentist.Apellido, dentist.Nombre, dentist.Matricula, id)

	_, err := s.DB.Exec(query)
	if err != nil {
		return dentist, err
	}

	/*lastID, err := result.LastInsertId()
	if lastID == 0 {
		return dentist, fmt.Errorf("dentista con ID %d no encontrado", id)
	}*/

	dentist.ID = int(id)
	return dentist, nil
}

func (s *SqlStore) Patch(id int, dentistPatch dentist.DentistPatch) (dentist.Dentist, error) {

	fieldstoUpdate, err := utils.StructToKeyValueString(dentistPatch)
	if err != nil {
		return dentist.Dentist{}, err
	}

	query := fmt.Sprintf("UPDATE dentista SET %s WHERE id = %d", fieldstoUpdate, id)

	_, err = s.DB.Exec(query)
	if err != nil {
		return dentist.Dentist{}, err
	}

	getDentist, err := s.GetByID(id)
	if err != nil {
		return getDentist, err
	}
	return getDentist, nil
}

func (s *SqlStore) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM dentista WHERE id = %d;", id)
	result, err := s.DB.Exec(query)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("borrado dentista con ID %d no encontrado", id)
	}

	return nil
}

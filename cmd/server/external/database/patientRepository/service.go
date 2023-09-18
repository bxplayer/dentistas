package database

import (
	"database/sql"
	"dentistas/internal/patient"
	"dentistas/internal/utils"
	"fmt"
)

type SqlStore struct {
	*sql.DB
}

func NewDatabase(db *sql.DB) *SqlStore {
	return &SqlStore{db}
}

func (s *SqlStore) GetByID(id int) (patient.Patient, error) {

	var p patient.Patient

	query := fmt.Sprintf("SELECT * FROM paciente WHERE id = %d;", id)
	row := s.DB.QueryRow(query)
	err := row.Scan(&p.ID, &p.Apellido, &p.Nombre, &p.Domicilio, &p.DNI, &p.Fecha_Alta)

	if err != nil {
		return patient.Patient{}, err
	}

	return p, nil

}

func (s *SqlStore) Create(patient patient.Patient) (patient.Patient, error) {

	query := fmt.Sprintf("INSERT INTO paciente (apellido, nombre, domicilio, dni, fecha_alta) VALUES ('%s', '%s', '%s', '%s', '%s');",
		patient.Apellido, patient.Nombre,patient.Domicilio, patient.DNI, patient.Fecha_Alta)

	result, err := s.DB.Exec(query)
	if err != nil {
		return patient, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return patient, err
	}

	patient.ID = int(id)
	return patient, nil
}

func (s *SqlStore) Update(id int, patient patient.Patient) (patient.Patient, error) {

	query := fmt.Sprintf("UPDATE paciente SET apellido = '%s', nombre = '%s', domicilio = '%s', DNI = '%s', fecha_alta = '%s' WHERE id = %d;",
		patient.Apellido, patient.Nombre, patient.Domicilio, patient.DNI, patient.Fecha_Alta, id)

	_, err := s.DB.Exec(query)
	if err != nil {
		return patient, err
	}

	patient.ID = int(id)
	return patient, nil
}

func (s *SqlStore) Patch(id int, patientPatch patient.PatientPatch) (patient.Patient, error) {

	fieldstoUpdate, err := utils.StructToKeyValueString(patientPatch)
	if err != nil {
		return patient.Patient{}, err
	}

	query := fmt.Sprintf("UPDATE paciente SET %s WHERE id = %d", fieldstoUpdate, id)

	_, err = s.DB.Exec(query)
	if err != nil {
		return patient.Patient{}, err
	}

	getPatient, err := s.GetByID(id)
	if err != nil {
		return getPatient, err
	}
	return getPatient, nil
}

func (s *SqlStore) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM paciente WHERE id = %d;", id)
	result, err := s.DB.Exec(query)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("borrado paciente con ID %d no encontrado", id)
	}

	return nil
}

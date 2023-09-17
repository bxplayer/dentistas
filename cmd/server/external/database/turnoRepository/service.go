package turnoRepository

import (
	"database/sql"
	"dentistas/internal/turno"
	"time"

	"fmt"
)

type SqlStore struct {
	*sql.DB
}

func NewDatabase(db *sql.DB) *SqlStore {
	return &SqlStore{db}
}

func (s *SqlStore) GetByID(id int) (turno.Turno, error) {

	var t turno.Turno

	query := fmt.Sprintf("SELECT * FROM turno WHERE id = %d;", id)
	row := s.DB.QueryRow(query)
	err := row.Scan(&t.ID, &t.Descripcion, &t.FechaHora, &t.PacienteId, &t.DentistaId)

	if err != nil {
		return turno.Turno{}, err
	}

	return t, nil
}

func (s *SqlStore) Create(turnoParam turno.Turno) (turno.Turno, error) {
	var date, err = time.Parse("2006-01-02 15:04:05", turnoParam.FechaHora)
	dateUTC := date.UTC().Format("2006-01-02 15:04:05")

	if err != nil {
		return turno.Turno{}, err
	}

	query := fmt.Sprintf("INSERT INTO turno (descripcion, fecha_hora, paciente_id, dentista_id) VALUES ('%v', '%v', %v, %v);", turnoParam.Descripcion, dateUTC, turnoParam.PacienteId, turnoParam.DentistaId)
	stmt, err := s.DB.Prepare(query)
	if err != nil {
		return turno.Turno{}, err
	}

	_, err = stmt.Exec()
	if err != nil {
		return turno.Turno{}, err
	}

	return turnoParam, nil

}

func (s *SqlStore) Update(id int, turno turno.Turno) (turno.Turno, error) {
	//TODO implement me
	panic("implement me Update")
}

func (s *SqlStore) Delete(id int) error {
	//TODO implement me
	panic("implement me Delete")
}

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
	if err != nil {
		return turno.Turno{}, err
	}
	dateUTC := date.UTC().Format("2006-01-02 15:04:05")

	query := fmt.Sprintf("INSERT INTO turno (descripcion, fecha_hora, paciente_dni, dentista_id) VALUES ('%v', '%v', %v, %v);", turnoParam.Descripcion, dateUTC, turnoParam.PacienteId, turnoParam.DentistaId)
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

func (s *SqlStore) Update(id int, turnoParam turno.Turno) (turno.Turno, error) {
	// Verificar si el turno existe.
	_, err1 := s.GetByID(id)
	if err1 != nil {
		// Si hay un error, eso significa que el turno no existe.
		return turno.Turno{}, err1
	}

	var date, err = time.Parse("2006-01-02 15:04:05", turnoParam.FechaHora)
	if err != nil {
		return turno.Turno{}, err
	}
	dateUTC := date.UTC().Format("2006-01-02 15:04:05")

	query := fmt.Sprintf("UPDATE turno SET descripcion = '%v', fecha_hora = '%v', paciente_dni = %v, dentista_id = %v WHERE id = %v;", turnoParam.Descripcion, dateUTC, turnoParam.PacienteId, turnoParam.DentistaId, id)
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

func (s *SqlStore) Delete(id int) error {
	// Verificar si el turno existe.
	_, err1 := s.GetByID(id)
	if err1 != nil {
		// Si hay un error, eso significa que el turno no existe.
		return err1
	}
	query := fmt.Sprintf("DELETE FROM turno WHERE id = %v;", id)
	stmt, err := s.DB.Prepare(query)
	if err != nil {
		return err
	}

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	return nil
}

func (s *SqlStore) GetByPacienteDNI(dni string) ([]turno.Turno, error) {
	var turnos []turno.Turno
	query := fmt.Sprintf("SELECT * FROM turno WHERE paciente_dni =  '%v';", dni)
	rows, err := s.DB.Query(query)
	if err != nil {
		return []turno.Turno{}, err
	}
	for rows.Next() {
		var t turno.Turno
		err := rows.Scan(&t.ID, &t.Descripcion, &t.FechaHora, &t.PacienteId, &t.DentistaId)
		if err != nil {
			return []turno.Turno{}, err
		}
		turnos = append(turnos, t)
	}
	return turnos, nil
}

func (s *SqlStore) SomeUpdate(id int, turnoParam turno.Turno) (turno.Turno, error) {
	// Verificar si el turno existe.
	savedTurno, err1 := s.GetByID(id)
	if err1 != nil {
		// Si hay un error, eso significa que el turno no existe.
		return turno.Turno{}, err1
	}

	if turnoParam.Descripcion != "" {
		savedTurno.Descripcion = turnoParam.Descripcion
	}
	if turnoParam.FechaHora != "" {
		var date, err = time.Parse("2006-01-02 15:04:05", turnoParam.FechaHora)
		if err != nil {
			return turno.Turno{}, err
		}
		dateUTC := date.UTC().Format("2006-01-02 15:04:05")
		savedTurno.FechaHora = dateUTC
	}
	if turnoParam.PacienteId != "" {
		savedTurno.PacienteId = turnoParam.PacienteId
	}
	if turnoParam.DentistaId != "" {
		savedTurno.DentistaId = turnoParam.DentistaId
	}

	query := fmt.Sprintf("UPDATE turno SET descripcion = '%v', fecha_hora = '%v', paciente_dni = %v, dentista_id = %v WHERE id = %v;", savedTurno.Descripcion, savedTurno.FechaHora, savedTurno.PacienteId, savedTurno.DentistaId, id)
	stmt, err := s.DB.Prepare(query)
	if err != nil {
		return turno.Turno{}, err
	}

	_, err = stmt.Exec()
	if err != nil {
		return turno.Turno{}, err
	}

	return savedTurno, nil
}

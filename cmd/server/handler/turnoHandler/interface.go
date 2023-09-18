package turnoHandler

import "dentistas/internal/turno"

type TurnoGetter interface {
	GetById(id int) (turno.Turno, error)
	GetByPacienteDNI(dni string) ([]turno.Turno, error)
}

type TurnoUpdater interface {
	Update(id int, turno turno.Turno) (turno.Turno, error)
	SomeUpdate(id int, turnoRepository turno.Turno) (turno.Turno, error)
}

type TurnoCreator interface {
	AddTurno(turno turno.Turno) (turno.Turno, error)
}

type TurnoDeleter interface {
	DeleteByID(id int) error
}

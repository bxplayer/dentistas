package turno

type Repository interface {
	GetByID(id int) (Turno, error)
	GetByPacienteDNI(dni string) ([]Turno, error)
	Update(id int, turno Turno) (Turno, error)
	Create(turno Turno) (Turno, error)
	Delete(id int) error
	//SomeUpdate(id int, turnoRepository Turno) (Turno, error)
}

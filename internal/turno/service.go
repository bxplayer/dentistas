package turno

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) GetById(id int) (Turno, error) {
	return s.repository.GetByID(id)
}

func (s *Service) Update(id int, turno Turno) (Turno, error) {
	return s.repository.Update(id, turno)
}

func (s *Service) AddTurno(turno Turno) (Turno, error) {
	return s.repository.Create(turno)
}

func (s *Service) DeleteByID(id int) error {
	return s.repository.Delete(id)
}

func (s *Service) GetByPacienteDNI(dni string) ([]Turno, error) {
	return s.repository.GetByPacienteDNI(dni)
}

//func (s *Service) SomeUpdate(id int, turnoRepository Turno) (Turno, error) {
//	return s.repository.SomeUpdate(id, turnoRepository)
//}

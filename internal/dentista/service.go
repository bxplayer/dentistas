package dentista

type Repository interface {
	GetByID(id int) (Dentista, error)
	Update(id int, dentista Dentista) (Dentista, error)
	Add(dentista Dentista) (Dentista, error)
	Delete(id int) error
}

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) GetByID(id int) (Dentista, error) {
	return s.repository.GetByID(id)
}

func (s *Service) ModifyByID(id int, dentista Dentista) (Dentista, error) {
	return s.repository.Update(id, dentista)
}

func (s *Service) AddDentista(dentista Dentista) (Dentista, error) {
	return s.repository.Add(dentista)
}

func (s *Service) DeleteByID(id int) error {
	return s.repository.Delete(id)
}
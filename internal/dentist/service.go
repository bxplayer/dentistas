package dentist

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) GetByID(id int) (Dentist, error) {
	return s.repository.GetByID(id)
}

func (s *Service) Update(id int, dentist Dentist) (Dentist, error) {
	return s.repository.Update(id, dentist)
}

func (s *Service) Add(dentist Dentist) (Dentist, error) {
	return s.repository.Add(dentist)
}

func (s *Service) Delete(id int) error {
	return s.repository.Delete(id)
}

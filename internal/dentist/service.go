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

func (s *Service) Patch(id int, dentist DentistPatch) (Dentist, error) {
	return s.repository.Patch(id, dentist)
}

func (s *Service) Create(dentist Dentist) (Dentist, error) {
	return s.repository.Create(dentist)
}

func (s *Service) Delete(id int) error {
	return s.repository.Delete(id)
}

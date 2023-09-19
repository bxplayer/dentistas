package patient

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) GetByID(id int) (Patient, error) {
	return s.repository.GetByID(id)
}

func (s *Service) Update(id int, patient Patient) (Patient, error) {
	return s.repository.Update(id, patient)
}

func (s *Service) Patch(id int, patient PatientPatch) (Patient, error) {
	return s.repository.Patch(id, patient)
}

func (s *Service) Create(patient Patient) (Patient, error) {
	return s.repository.Create(patient)
}

func (s *Service) Delete(id int) error {
	return s.repository.Delete(id)
}

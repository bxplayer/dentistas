package patient

type Repository interface {
	Create(patient Patient) (Patient, error)
	GetByID(id int) (Patient, error)
	Update(id int, patient Patient) (Patient, error)
	Patch(id int, patient PatientPatch) (Patient, error)
	Delete(id int) error
}

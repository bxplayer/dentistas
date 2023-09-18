package dentist

type Repository interface {
	Create(dentist Dentist) (Dentist, error)
	GetByID(id int) (Dentist, error)
	Update(id int, dentist Dentist) (Dentist, error)
	Patch(id int, dentist DentistPatch) (Dentist, error)
	Delete(id int) error
}

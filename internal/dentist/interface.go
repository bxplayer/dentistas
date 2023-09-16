package dentist

type Repository interface {
	Add(dentist Dentist) (Dentist, error)
	GetByID(id int) (Dentist, error)
	Update(id int, dentist Dentist) (Dentist, error)
	Delete(id int) error
}

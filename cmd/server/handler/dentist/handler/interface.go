package handler

import "dentistas/internal/dentist"

type DentistGetters interface {
	GetByID(id int) (dentist.Dentist, error)
}

type DentistCreator interface {
	Create(dentist dentist.Dentist) (dentist.Dentist, error)
}

type DentistUpdates interface {
	Update(id int, dentist dentist.Dentist) (dentist.Dentist, error)
	Patch(id int, dentist dentist.DentistPatch) (dentist.Dentist, error)
}

type DentistDelete interface {
	Delete(id int) error
}

package handler

import "dentistas/internal/dentist"

type DentistGetters interface {
	GetByID(id int) (dentist.Dentist, error)
}

type DentistCreator interface {
	Add(dentista dentist.Dentist) (dentist.Dentist, error)
}

type DentistUpdates interface {
	Update(id int, dentista dentist.Dentist) (dentist.Dentist, error)
}

type DentistDelete interface {
	Delete(id int) error
}

package patientHandler

import "dentistas/internal/patient"

type PatientGetters interface {
	GetByID(id int) (patient.Patient, error)
}

type PatientCreator interface {
	Create(patient patient.Patient) (patient.Patient, error)
}

type PatientUpdates interface {
	Update(id int, patient patient.Patient) (patient.Patient, error)
	Patch(id int, patient patient.PatientPatch) (patient.Patient, error)
}

type PatientDelete interface {
	Delete(id int) error
}

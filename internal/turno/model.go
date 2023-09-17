package turno

type Turno struct {
	ID          int    `json:"id"`
	Descripcion string `json:"descripcion"`
	FechaHora   string `json:"fechaHora"`
	PacienteId  string `json:"pacienteId"`
	DentistaId  string `json:"dentistaId"`
}

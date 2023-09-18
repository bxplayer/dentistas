package turno

type Turno struct {
	ID          int    `json:"id"`
	Descripcion string `json:"descripcion"`
	FechaHora   string `json:"fechaHora"`
	PacienteDni string `json:"pacienteDni"`
	DentistaId  string `json:"dentistaId"`
}

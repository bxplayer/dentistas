package patient

type Patient struct {
	ID        int    `json:"id"`
	Apellido  string `json:"apellido" binding:"required"`
	Nombre    string `json:"nombre" binding:"required"`
	Domicilio string `json:"domicilio" binding:"required"`
	DNI string `json:"dni"  binding:"required"`
	Fecha_Alta string `json:"fecha_alta" binding:"required"`
}

type PatientPatch struct {
	Apellido  string `json:"apellido"`
	Nombre    string `json:"nombre"`
	Domicilio string `json:"domicilio"`
	DNI string `json:"dni"  binding:"required"`
	Fecha_Alta string `json:"fecha_alta"`
}

package paciente

type Paciente struct {
	ID          int    `json:"id"`
	Nombre      string `json:"nombre"`
	Apellido    string `json:"apellido"`
	Domicilio   string `json:"domicilio"`
	DNI         string `json:"dni"`
	FechaDeAlta string `json:"fecha_de_alta"`
}

package paciente

import "time"

type Paciente struct {
	ID          int       `json:"id"`
	Nombre      string    `json:"nombre"`
	Apellido    string    `json:"apellido"`
	Domicilio   string    `json:"domicilio"`
	DNI         string    `json:"dni"`
	FechaDeAlta time.Time `json:"fecha_de_alta"`
}

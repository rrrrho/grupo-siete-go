package turno

import (
	"grupo-siete-go/internal/odontologo"
	"grupo-siete-go/internal/paciente"
)

type Turno struct {
	ID          int                   `json:"id"`
	Paciente    paciente.Paciente     `json:"paciente"`
	Odontologo  odontologo.Odontologo `json:"odontologo"`
	FechaYHora  string                `json:"fecha_y_hora"`
	Descripcion string                `json:"descripcion"`
}

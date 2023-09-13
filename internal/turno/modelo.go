package turno

import (
	"grupo-siete-go/internal/odontologo"
	"grupo-siete-go/internal/paciente"
	"time"
)

type Turno struct {
	Paciente    paciente.Paciente     `json:"paciente"`
	Odontologo  odontologo.Odontologo `json:"odontologo"`
	FechaYHora  time.Time             `json:"fecha_y_hora"`
	Descripcion string                `json:"descripcion"`
}

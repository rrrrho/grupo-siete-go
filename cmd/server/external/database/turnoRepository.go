package database

import (
	"database/sql"
	"errors"
	"fmt"
	modelo "grupo-siete-go/internal/turno"
)

type SqlStore struct {
	*sql.DB
}

func NewDatabase(db *sql.DB) *SqlStore {
	return &SqlStore{db}
}

func (s *SqlStore) Save(turno modelo.Turno) (modelo.Turno, error) {
	// valido que exista el paciente
	_, err := s.DB.Query(fmt.Sprintf("SELECT * FROM pacientes WHERE id = %d;", turno.Paciente.ID))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			fmt.Println("El paciente no existe en la base de datos:", err)
		} else {
			fmt.Println("Error al ejecutar la consulta:", err)
		}
		return modelo.Turno{}, err
	}

	// valido que exista el odontologo
	_, err = s.DB.Query(fmt.Sprintf("SELECT * FROM odontologos WHERE id = %d;", turno.Odontologo.ID))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			fmt.Println("El odontologo no existe en la base de datos:", err)
		} else {
			fmt.Println("Error al ejecutar la consulta:", err)
		}
		return modelo.Turno{}, err
	}

	// guardo el turno
	res, err := s.DB.Exec("INSERT INTO turnos (paciente_id, odontologo_id, fecha_hora, descripcion) VALUES (%d, %d, %s, %s)",
		turno.Paciente.ID, turno.Odontologo.ID, turno.FechaYHora, turno.Descripcion)
	if err != nil {
		fmt.Println("Error al ejecutar la consulta:", err)
		return modelo.Turno{}, err
	}

	// obtengo el ID del turno guardado y lo seteo al turno del parametro para devolverlo
	var aux int64
	aux, err = res.LastInsertId()
	if err != nil {
		fmt.Println("Error al ejecutar la consulta:", err)
		return modelo.Turno{}, err
	}
	turno.ID = int(aux)

	return turno, nil
}

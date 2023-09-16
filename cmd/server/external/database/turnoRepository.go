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
	res, err := s.DB.Exec("INSERT INTO turnos (paciente_id, odontologo_id, fecha_hora, descripcion) VALUES (%d, %d, %s, %s);",
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

func (s *SqlStore) GetByID(id int) (modelo.Turno, error) {
	var turno modelo.Turno

	// obtengo el turno
	res := s.DB.QueryRow("SELECT t.id, p.id, p.nombre, p.apellido, p.domicilio, p.dni, p.fecha_de_alta, "+
		"o.id, o.nombre, o.apellido, o.matricula, t.fecha_hora, t.descripcion "+
		"FROM turnos t INNER JOIN pacientes p ON t.paciente_id = p.id "+
		"INNER JOIN odontologos o ON t.odontologo_id = o.id WHERE id = ?", id)
	err := res.Scan(&turno.ID, &turno.Paciente.ID, &turno.Paciente.Nombre, &turno.Paciente.Apellido, &turno.Paciente.Domicilio, &turno.Paciente.DNI,
		&turno.Paciente.FechaDeAlta, &turno.Odontologo.ID, &turno.Odontologo.Nombre, &turno.Odontologo.Apellido, &turno.Odontologo.Matricula,
		&turno.FechaYHora, &turno.Descripcion)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			fmt.Println("El turno no existe en la base de datos:", err)
		} else {
			fmt.Println("Error al ejecutar la consulta:", err)
		}
		return modelo.Turno{}, err
	}

	return turno, nil
}

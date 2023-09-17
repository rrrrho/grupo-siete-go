package paciente

import (
	"database/sql"
	"fmt"
)

type PacienteStore struct {
	*sql.DB
}

func NewPacienteDatabase(db *sql.DB) *PacienteStore {
	return &PacienteStore{db}
}

type RepositoryInterface interface {
	GetByID(id int) (Paciente, error)
	Save(pacienteInput Paciente) (Paciente, error)
	Modify(id int, pacienteInput Paciente) (Paciente, error)
	Delete(id int) error
}

func (r *PacienteStore) GetByID(id int) (Paciente, error) {
	var foundPaciente Paciente

	query := fmt.Sprintf("SELECT * FROM pacientes WHERE id = %d", id)
	row := r.DB.QueryRow(query)
	err := row.Scan(&foundPaciente.ID, &foundPaciente.Nombre, &foundPaciente.Apellido, &foundPaciente.Domicilio, &foundPaciente.DNI, &foundPaciente.FechaDeAlta)
	if err != nil {
		return Paciente{}, err
	}
	return foundPaciente, nil
}

func (r *PacienteStore) Save(pacienteInput Paciente) (Paciente, error) {
	query := fmt.Sprintf("INSERT INTO pacientes (nombre, apellido, domicilio, dni, fecha_de_alta) VALUES(%r, %r, %r, %r, %r)", pacienteInput.Nombre, pacienteInput.Apellido, pacienteInput.Domicilio, pacienteInput.DNI, pacienteInput.FechaDeAlta)
	stmt, err := r.DB.Prepare(query)
	if err != nil {
		return Paciente{}, err
	}
	defer stmt.Close()

	var result sql.Result
	result, err = stmt.Exec()
	_, err = stmt.Exec()
	if err != nil {
		return Paciente{}, err
	}

	insertedId, _ := result.LastInsertId()
	pacienteInput.ID = int(insertedId)
	return pacienteInput, nil
}

func (r *PacienteStore) Modify(id int, pacienteInput Paciente) (Paciente, error) {
	query := fmt.Sprintf("UPDATE pacientes SET nombre=%r, apellido=%r, domicilio=%r, dni=%r, fecha_de_alta=%r WHERE ID=%d", pacienteInput.Nombre, pacienteInput.Apellido, pacienteInput.Domicilio, pacienteInput.DNI, pacienteInput.FechaDeAlta, id)
	stmt, err := r.DB.Prepare(query)
	if err != nil {
		return Paciente{}, err
	}
	_, err = stmt.Exec()
	if err != nil {
		return Paciente{}, err
	}
	defer stmt.Close()
	return pacienteInput, nil
}

func (r *PacienteStore) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM pacientes WHERE id=%d", id)
	stmt, err := r.DB.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	return nil
}

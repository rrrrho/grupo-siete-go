package paciente

import (
	"database/sql"
	"fmt"
)

type Repository struct {
	*sql.DB
}

func NewDatabase(db *sql.DB) *Repository {
	return &Repository{db}
}

type RepositoryInterface interface {
	GetByID(id int) (Paciente, error)
	Save(pacienteInput Paciente) (Paciente, error)
	Modify(id int, pacienteInput Paciente) (Paciente, error)
	Delete(id int) error
}

func (r *Repository) GetByID(id int) (Paciente, error) {
	var foundPaciente Paciente

	query := fmt.Sprintf("SELECT * FROM pacientes WHERE id = %d", id)
	row := r.DB.QueryRow(query)
	err := row.Scan(&foundPaciente.ID, &foundPaciente.Nombre, &foundPaciente.Apellido, &foundPaciente.Domicilio, &foundPaciente.DNI, &foundPaciente.FechaDeAlta)
	if err != nil {
		return Paciente{}, err
	}
	return foundPaciente, nil
}

func (r *Repository) Save(pacienteInput Paciente) (Paciente, error) {
	query := "INSERT INTO pacientes (nombre, apellido, domicilio, dni, fecha_de_alta) VALUES (?, ?, ?, ?, ?)"
    stmt, err := r.DB.Prepare(query)
    if err != nil {
        return Paciente{}, err
    }
    defer stmt.Close()
    fechaFormatted := pacienteInput.FechaDeAlta.Format("2006-01-02")

    result, err := stmt.Exec(
        pacienteInput.Nombre,
        pacienteInput.Apellido,
        pacienteInput.Domicilio,
        pacienteInput.DNI,
        fechaFormatted,
    )
    if err != nil {
        return Paciente{}, err
    }

    insertedID, err := result.LastInsertId()
    if err != nil {
        return Paciente{}, err
    }

    pacienteInput.ID = int(insertedID)
    return pacienteInput, nil
}

func (r *Repository) Modify(id int, pacienteInput Paciente) (Paciente, error) {
	query := "UPDATE pacientes SET nombre=?, apellido=?, domicilio=?, dni=?, fecha_de_alta=? WHERE ID=?"
    stmt, err := r.DB.Prepare(query)
    if err != nil {
        return Paciente{}, err
    }
    defer stmt.Close()

    fechaFormatted := pacienteInput.FechaDeAlta.Format("2006-01-02")

    _, err = stmt.Exec(
        pacienteInput.Nombre,
        pacienteInput.Apellido,
        pacienteInput.Domicilio,
        pacienteInput.DNI,
        fechaFormatted,
        id,
    )
    if err != nil {
        return Paciente{}, err
    }

    return pacienteInput, nil
}

func (r *Repository) Delete(id int) error {
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

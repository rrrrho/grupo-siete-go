package database

import (
	"database/sql"
	"fmt"
	"grupo-siete-go/internal/odontologo"
)

type SqlStore struct {
	*sql.DB 
}

func NewDatabase(db *sql.DB) *SqlStore {
	return &SqlStore{db}
}

func (s *SqlStore) GetByID(id int) (odontologo.Odontologo, error) {
	var foundOdontologo odontologo.Odontologo

	query := fmt.Sprintf("SELECT * FROM odontologos WHERE id = %d", id)
	row := s.DB.QueryRow(query)
	err := row.Scan(&foundOdontologo.ID, &foundOdontologo.Nombre, &foundOdontologo.Apellido, &foundOdontologo.Matricula)
	if err != nil {
		return odontologo.Odontologo{}, err
	}
	return foundOdontologo, nil
}

func (s *SqlStore) Modify(id int, odontologoInput odontologo.Odontologo) (odontologo.Odontologo, error) {
	query := fmt.Sprintf("UPDATE odontologos SET nombre=%s, apellido=%s, matricula=%s WHERE ID=%d", odontologoInput.Nombre, odontologoInput.Apellido, odontologoInput.Matricula, odontologoInput.ID)
	stmt, err := s.DB.Prepare(query)
	if err != nil {
		return odontologo.Odontologo{}, err
	}
	_, err = stmt.Exec()
	if err != nil {
		return odontologo.Odontologo{}, err
	}
	defer stmt.Close()
	return odontologoInput, nil
}

func (s *SqlStore) Save(odontologoInput odontologo.Odontologo) (odontologo.Odontologo, error) {
	query := fmt.Sprintf("INSERT INTO odontologos (nombre, apellido, matricula) VALUES(%s, %s, %s)", odontologoInput.Nombre, odontologoInput.Apellido, odontologoInput.Matricula)
	stmt, err := s.DB.Prepare(query)
	if err != nil {
		return odontologo.Odontologo{}, err
	}
	defer stmt.Close()

	var result sql.Result
	result, err = stmt.Exec()
	_, err = stmt.Exec()
	if err != nil {
		return odontologo.Odontologo{}, err
	}

	insertedId, _ := result.LastInsertId()
	odontologoInput.ID = int(insertedId)
	return odontologoInput, nil
}

func (s *SqlStore) Delete(id int) (string, error) {
	resultString := ""
	query := fmt.Sprintf("DELETE FROM odontologos WHERE id=%d", id)
	stmt, err := s.DB.Prepare(query)
	if err != nil {
		return resultString, err
	}
	_, err = stmt.Exec()
	if err != nil {
		return resultString, err
	}
	return fmt.Sprintf("Odontologo ID %d elminado correctamente", id), nil
}

package database

import (
	"database/sql"
	"fmt"
	"grupo-siete-go/internal/odontologo"
	"errors"
)

type OdontologoStore struct {
	*sql.DB
}

func NewOdontologoDatabase(db *sql.DB) *OdontologoStore {
	return &OdontologoStore{db}
}

func (s *OdontologoStore) GetByID(id int) (odontologo.Odontologo, error) {
	var foundOdontologo odontologo.Odontologo

	query := fmt.Sprintf("SELECT * FROM odontologos WHERE id = %d", id)
	row := s.DB.QueryRow(query)
	err := row.Scan(&foundOdontologo.ID, &foundOdontologo.Nombre, &foundOdontologo.Apellido, &foundOdontologo.Matricula)
	if err != nil {
		return odontologo.Odontologo{}, err
	}
	return foundOdontologo, nil
}


func (s *OdontologoStore) Update(id int, odontologoInput odontologo.Odontologo) (odontologo.Odontologo, error) {
	// valido que exista el Odontologo
	_, err := s.DB.Query("SELECT * FROM odontologos WHERE id = ?", id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			fmt.Println("El Odontologo no existe en la base de datos:", err)
		} else {
			fmt.Println("Error al ejecutar la consulta:", err)
		}
		return odontologo.Odontologo{}, err
	}

	// armo la query
	query := "UPDATE odontologos SET"

	if odontologoInput.ID > 0 {
		query += " id = '" + string(odontologoInput.ID) + "',"
	}
	if odontologoInput.Nombre != "" {
		query += " nombre = '" + odontologoInput.Nombre + "',"
	}
	if odontologoInput.Apellido != "" {
		query += " apellido = '" + odontologoInput.Apellido + "',"
	}
	if odontologoInput.Matricula != "" {
		query += " matricula = '" + odontologoInput.Matricula + "',"
	}

	query = query[0 : len(query)-1]
	query += " WHERE id = ?"

	fmt.Println(query)
	// actualizo el odontologo
	_, err = s.DB.Exec(query, id)

	if err != nil {
		fmt.Println("Error al ejecutar la consulta:", err)
		return odontologo.Odontologo{}, err
	}

	// obtengo el recurso actualizado
	updatedOdontologo, err := s.GetByID(id)
	return updatedOdontologo, nil
}

func (s *OdontologoStore) Replace(odontologoInput odontologo.Odontologo) (odontologo.Odontologo, error) {
	_, err := s.DB.Exec("UPDATE odontologos SET nombre=?, apellido=?, matricula=? WHERE ID=?;", odontologoInput.Nombre, odontologoInput.Apellido, odontologoInput.Matricula, odontologoInput.ID)
	if err != nil {
		fmt.Println("Error al ejecutar la consulta:", err)
		return odontologo.Odontologo{}, err
	}
	updatedOdontologo, err := s.GetByID(odontologoInput.ID)
	if err != nil {
		fmt.Println("Error al ejecutar la consulta:", err)
		return odontologo.Odontologo{}, err
	}
	return updatedOdontologo, nil
}

func (s *OdontologoStore) Save(odontologoInput odontologo.Odontologo) (odontologo.Odontologo, error) {
	res, err := s.DB.Exec("INSERT INTO odontologos (nombre, apellido, matricula) VALUES(?,?,?);", odontologoInput.Nombre, odontologoInput.Apellido, odontologoInput.Matricula)
	if err != nil {
	fmt.Println("Error al ejecutar la consulta:", err)
	return odontologo.Odontologo{}, err
	}
	insertedId, _ := res.LastInsertId()
	odontologoInput.ID = int(insertedId)
	return odontologoInput, nil
}

func (s *OdontologoStore) Delete(id int) (string, error) {
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

package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func NewMySQLDatabase(host, port, user, password, dbname string) (*sql.DB, error) {
	connStr := "host=" + host + " port=" + port + " user=" + user + " password=" + password + " dbname=" + dbname + " sslmode=disable"

	db, err := sql.Open("mysql", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

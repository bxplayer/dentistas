package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func NewMysqlDatabase(host, port, user, password, dbname string) (*sql.DB, error) {

	connStr := fmt.Sprintf("%s:%s@(%s:%s)/%s", user, password, host, port, dbname)

	db, err := sql.Open("mysql", connStr)
	if err != nil {
		return nil, err
	}

	// Probamos la conexión para asegurarnos de que sea válida.
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil

}

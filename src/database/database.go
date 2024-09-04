package database

import (
	"api/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// MakeConnection make a connection with databse
func MakeConnection() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.UrlDatabaseConnection)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

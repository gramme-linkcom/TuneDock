package db

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

func Open(path string) (*sql.DB, error) {
	database, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}

	if err := database.Ping(); err != nil {
		database.Close()
		return nil, err
	}

	if _, err := database.Exec(`PRAGMA foreign_keys = ON;`); err != nil {
		database.Close()
		return nil, err
	}

	return database, nil
}

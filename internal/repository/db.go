package repository

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func InitDb() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./jobportal.db")
	if err != nil {
		return nil, err
	}

	// Check if connection is valid
	if err = db.Ping(); err != nil {
		return nil, err
	}

	// Create table
	if err = createTable(db); err != nil {
		return nil, err
	}

	return db, nil
}

func createTable(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS users(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		is_admin BOOLEAN DEFAULT 0,
		profile_picture TEXT
	)`)
	return err
}

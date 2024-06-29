package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func NewPGConnection(connString string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	fmt.Println("Connected to PostgreSQL database!")
	return db, nil
}

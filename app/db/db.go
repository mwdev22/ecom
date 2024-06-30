package dbConn

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func DbOpen(connString string, database string) (*sql.DB, error) {
	db, err := sql.Open(database, connString)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to PostgreSQL database!")
	return db, nil
}

func InitConn(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		db.Close()
		log.Fatal(err)
	}
}

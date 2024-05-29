package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func MustConnect(connectionString string) (*sql.DB, func()) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Could not ping database: %v", err)
	}

	cleanup := func() {
		err := db.Close()
		if err != nil {
			log.Printf("Error closing database: %v", err)
		}
	}

	return db, cleanup
}

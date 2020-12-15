package repository

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // For use in sql import
)

//CockroachRepository is the concrete implementation of IRepository using CockroachDB
type CockroachRepository struct{}

// InsertPlayer Inserts a newly registered player into the database
func (r CockroachRepository) InsertPlayer() {
	// Connect
	var connectionString string = "postgresql://root@localhost:26257/snake_fever_db?sslmode=disable"

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}

	// Insert
	if _, err := db.Exec(`INSERT INTO tbl_player (created_at, updated_at) VALUES (NOW(), NOW());`); err != nil {
		log.Fatal(err)
	}
}

package repository

import (
	"database/sql"
	"log"

	"snake-fever/snake-fever/pkg/model"

	_ "github.com/lib/pq" // For use in sql import
)

//CockroachRepository is the concrete implementation of IRepository using CockroachDB
type CockroachRepository struct{}

// InsertPlayer is a method for CockroachRepository that inserts a newly registered player into the databasev
func (r CockroachRepository) InsertPlayer(playerObject model.Player) {
	// Connect
	var connectionString string = "postgresql://root@localhost:26257/snake_fever_db?sslmode=disable"

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}

	// Insert
	if _, err := db.Query(`INSERT INTO tbl_player (username, created_at, updated_at) VALUES ($1, NOW(), NOW());`, playerObject.Username); err != nil {
		log.Fatal(err)
	}
}

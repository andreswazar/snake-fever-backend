package cockroachdb

import (
	"database/sql"
	"log"
	"snake-fever/snake-fever/pkg/model"

	_ "github.com/lib/pq" // For use in sql import
)

//PlayerRepository is the concrete implementation of IPlayerRepository using CockroachDB
type PlayerRepository struct{}

func connectToDB() *sql.DB {
	var connectionString string = "postgresql://root@localhost:26257/snake_fever_db?sslmode=disable"

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}

	return db
}

// InsertPlayer is a method for PlayerRepository that inserts a newly registered player into the database
func (r PlayerRepository) InsertPlayer(playerObject model.Player) {
	db := connectToDB()
	defer db.Close()

	// Insert
	if _, err := db.Query(`INSERT INTO tbl_player (username, created_at, updated_at) VALUES ($1, NOW(), NOW());`, playerObject.Username); err != nil {
		log.Fatal(err)
	}
}

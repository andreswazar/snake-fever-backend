package cockroachdb

import (
	"fmt"
	"snake-fever/snake-fever/pkg/model"

	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq" // For use in sql import
)

//ScoreRepository is the concrete implementation of IScoreRepository using CockroachDB
type ScoreRepository struct{}

func connectToDB() *sql.DB {
	var connectionString string = "postgresql://root@localhost:26257/snake_fever_db?sslmode=disable"

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}

	return db
}

// InsertScore is a method for ScoreRepository that inserts a newly registered score into the database
func (r ScoreRepository) InsertScore(scoreObject model.Score) error {
	db := connectToDB()
	defer db.Close()

	// Create
	_, err := db.Query(`INSERT INTO tbl_score (score, player, created_at) VALUES ($1, $2, NOW());`, scoreObject.PointsScored, scoreObject.PlayerUsername)
	if err != nil {
		fmt.Printf("%v\n", err)
		return err
	}

	return nil
}

// GetAllScores is a method for ScoreRepository that queries all the entries in the scores table
func (r ScoreRepository) GetAllScores() ([]model.Score, error) {
	db := connectToDB()
	defer db.Close()

	result := make([]model.Score, 0, 5)

	// Read all
	rows, err := db.Query("SELECT * FROM tbl_score ORDER BY score DESC LIMIT 10")
	if err != nil {
		fmt.Printf("%v\n", err)
		return []model.Score{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var scoreID string
		var pointsScored int
		var playerUsername string
		var createdAt time.Time

		err := rows.Scan(&scoreID, &pointsScored, &playerUsername, &createdAt)
		if err != nil {
			fmt.Printf("%v", err)
			return []model.Score{}, err
		}

		scoreObject := model.Score{ID: scoreID, PointsScored: pointsScored, PlayerUsername: playerUsername, CreatedAt: createdAt}
		result = append(result, scoreObject)
	}

	return result, nil
}

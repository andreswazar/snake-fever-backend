package model

import "time"

// Score is a model for player scores
type Score struct {
	ID             string
	PointsScored   int
	PlayerUsername string
	CreatedAt      time.Time
}

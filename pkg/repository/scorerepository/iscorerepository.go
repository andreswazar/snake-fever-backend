package scorerepository

import (
	"snake-fever/snake-fever/pkg/model"
)

// IScoreRepository is an abstract implementation for repository
type IScoreRepository interface {
	InsertScore(scoreObject model.Score) error
	GetAllScores() ([]model.Score, error)
}

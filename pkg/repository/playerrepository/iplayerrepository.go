package playerrepository

import (
	"snake-fever/snake-fever/pkg/model"
)

// IPlayerRepository is an abstract implementation for repository
type IPlayerRepository interface {
	InsertPlayer(playerObject model.Player)
}

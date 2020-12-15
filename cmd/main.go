package main

import (
	"snake-fever/snake-fever/pkg/model"
	"snake-fever/snake-fever/pkg/repository"
)

// Dependency injection template
type dependency struct {
	playerRepository repository.IPlayerRepository
}

func main() {
	// Inject dependencies
	dependencies := dependency{
		playerRepository: new(repository.CockroachRepository),
	}

	playerObject := model.Player{Username: "eeeeeeeee"}
	dependencies.playerRepository.InsertPlayer(playerObject)
}

// To-do: insert players structs

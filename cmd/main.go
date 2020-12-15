package main

import (
	"snake-fever/snake-fever/pkg/repository"
)

func main() {
	var repos repository.IRepository = new(repository.CockroachRepository)
	repos.InsertPlayer()
}

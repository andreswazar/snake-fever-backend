package main

import (
	"snake-fever/snake-fever/pkg/repository/scorerepository"
	"snake-fever/snake-fever/pkg/repository/scorerepository/cockroachdb"

	"snake-fever/snake-fever/pkg/router"
	"snake-fever/snake-fever/pkg/router/chirouter"
)

// Dependency injection template
type program struct {
	scoreRepository scorerepository.IScoreRepository
	router          router.IRouter
}

// Inject dependencies
func (p *program) programConstructor(sr scorerepository.IScoreRepository, r router.IRouter) {
	p.scoreRepository = sr
	p.router = r
}

func main() {
	// Create objects with injected dependencies
	program := program{}
	program.programConstructor(new(cockroachdb.ScoreRepository), new(chirouter.Router))
	// Distribute dependencies
	program.router.RouterConstructor(program.scoreRepository)

	program.router.StartServer()
}

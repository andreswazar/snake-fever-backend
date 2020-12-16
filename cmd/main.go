package main

import (
	"encoding/json"
	"log"
	"net/http"

	"snake-fever/snake-fever/pkg/repository/playerrepository"
	"snake-fever/snake-fever/pkg/repository/playerrepository/cockroachdb"

	"snake-fever/snake-fever/pkg/router"
	"snake-fever/snake-fever/pkg/router/chirouter"
)

// Dependency injection template
type program struct {
	playerRepository playerrepository.IPlayerRepository
	router           router.IRouter
}

// Inject dependencies
func (p *program) startProgram() {
	p.playerRepository = new(cockroachdb.PlayerRepository)
	p.router = new(chirouter.Router)
}

func getRequestHandler(response http.ResponseWriter, request *http.Request) {
	json.NewEncoder(response).Encode("You hit the GET endpoint")
}

func main() {
	// Create object with injected dependencies
	program := new(program)
	program.startProgram()

	router := program.router.StartServer()
	router.Get("/api/getExample", getRequestHandler)
	log.Fatal(http.ListenAndServe(":8081", router))

	// playerObject := model.Player{Username: "Test 3"}
	// program.playerRepository.InsertPlayer(playerObject)
}

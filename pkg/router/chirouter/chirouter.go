package chirouter

import (
	"snake-fever/snake-fever/pkg/model"
	"snake-fever/snake-fever/pkg/repository/scorerepository"

	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

//Router is the concrete implementation of IRouter using Chi Router
type Router struct {
	scoreRepository scorerepository.IScoreRepository
}

//RouterConstructor is the dependency injector for the Router struct
func (r *Router) RouterConstructor(sr scorerepository.IScoreRepository) {
	r.scoreRepository = sr
}

// Handlers
// Returns a GET handler that has access to scoreRepository using closures
func (r Router) getAllScoresRequestHandler() http.HandlerFunc {
	scoreRepository := r.scoreRepository

	return func(response http.ResponseWriter, request *http.Request) {
		response.Header().Set("Content-Type", "application/json")

		scoresArray, err := scoreRepository.GetAllScores()
		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)
		} else {
			json.NewEncoder(response).Encode(scoresArray)
		}
	}
}

// Returns a POST handler that has access to scoreRepositoryu using closures
func (r Router) postScoreRequestHandler() http.HandlerFunc {
	scoreRepository := r.scoreRepository

	return func(response http.ResponseWriter, request *http.Request) {
		response.Header().Set("Content-Type", "application/json")

		// Add the request body to a new map
		mappedRequest := map[string]string{}
		json.NewDecoder(request.Body).Decode(&mappedRequest)

		// Add every entry of the map to the model.Score object and then create it on the database
		convertedScore, _ := strconv.Atoi(mappedRequest["pointsScored"])
		scoreObject := model.Score{PointsScored: convertedScore, PlayerUsername: mappedRequest["playerUsername"]}

		err := scoreRepository.InsertScore(scoreObject)
		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)
		} else {
			response.WriteHeader(http.StatusOK)
		}
	}
}

// StartServer is a method for Router that allows our server to start listening to requests
func (r Router) StartServer() {
	fmt.Println("Starting Server...")
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// Register handlers
	router.Get("/api/getAllScores", r.getAllScoresRequestHandler())
	router.Post("/api/postScore", r.postScoreRequestHandler())

	fmt.Println("Server started in port 8082")
	log.Fatal(http.ListenAndServe(":8082", router))
}

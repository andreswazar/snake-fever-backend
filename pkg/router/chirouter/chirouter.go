package chirouter

import (
	"fmt"

	"github.com/go-chi/chi"
)

//Router is the concrete implementation of IRouter using Chi Router
type Router struct{}

// StartServer is a method for Router that allows our server to start listening to requests
func (r Router) StartServer() *chi.Mux {
	fmt.Println("Starting Server...")
	router := chi.NewRouter()

	return router
}

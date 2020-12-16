package router

import "github.com/go-chi/chi"

// IRouter is an abstract implementation for router
type IRouter interface {
	StartServer() *chi.Mux
}

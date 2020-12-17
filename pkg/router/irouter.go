package router

import "snake-fever/snake-fever/pkg/repository/scorerepository"

// IRouter is an abstract implementation for router
type IRouter interface {
	RouterConstructor(sr scorerepository.IScoreRepository)
	StartServer()
}

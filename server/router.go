package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func SetupRouter(logger *zap.Logger, app SwapApp) http.Handler {
	r := &router{
		app:    app,
		logger: logger,
	}
	return r.Handler()
}

type SwapApp interface {
}

type router struct {
	app    SwapApp
	logger *zap.Logger
}

func (ro *router) Handler() http.Handler {
	r := mux.NewRouter()
	return r
}

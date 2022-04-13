package server

import (
	"net/http"

	"github.com/gorilla/handlers"
	"go.uber.org/zap"
)

type Server struct {
	addr   string
	app    SwapApp
	logger *zap.Logger
}

func NewServer(logger *zap.Logger, addr string, app SwapApp) *Server {
	return &Server{
		addr:   addr,
		app:    app,
		logger: logger,
	}
}

func (s *Server) Start() {
	go func() {
		headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
		originsOk := handlers.AllowedOrigins([]string{"*"})
		methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
		err := http.ListenAndServe(s.addr, handlers.CORS(originsOk, headersOk, methodsOk)(SetupRouter(s.logger, s.app)))
		if err != nil {
			s.logger.Fatal(err.Error())
		}
	}()
}

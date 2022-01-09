package http

import (
	"github.com/gin-gonic/gin"
	"go-backend-template/internal/usecase"
)

type Server struct {
	config   Config
	engine   *gin.Engine
	usecases *usecase.Usecases
}

func NewServer(config Config, usecases *usecase.Usecases) *Server {
	server := &Server{
		config:   config,
		engine:   gin.Default(),
		usecases: usecases,
	}

	router := newRouter(server)
	router.init()

	return server
}

func (s Server) Listen() error {
	return s.engine.Run(s.config.Address())
}

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
	gin.SetMode(gin.ReleaseMode)

	server := &Server{
		config:   config,
		engine:   gin.New(),
		usecases: usecases,
	}

	router := newRouter(server)
	router.init()

	return server
}

func (s Server) Listen() error {
	return s.engine.Run(s.config.Address())
}

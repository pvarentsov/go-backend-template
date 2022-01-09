package http

import (
	"fmt"
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
	fmt.Printf("API server listening at: %s\n\n", s.config.Address())
	return s.engine.Run(s.config.Address())
}

package http

import (
	"fmt"

	"go-backend-template/internal/usecase"
	"go-backend-template/internal/util/crypto"

	"github.com/gin-gonic/gin"
)

type Server struct {
	config   Config
	crypto   crypto.Crypto
	engine   *gin.Engine
	usecases *usecase.Usecases
}

func NewServer(config Config, crypto crypto.Crypto, usecases *usecase.Usecases) *Server {
	gin.SetMode(gin.ReleaseMode)

	server := &Server{
		config:   config,
		crypto:   crypto,
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

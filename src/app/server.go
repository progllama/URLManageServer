package app

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	engine *gin.Engine
}

func NewServer() *Server {
	return &Server{
		engine: gin.Default(),
	}
}

func (s *Server) Open() error {
	err := s.engine.Run("localhost:8000")
	return err
}

func (s *Server) Engine() *gin.Engine {
	return s.engine
}

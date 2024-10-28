package server

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	engine *gin.Engine
}

func NewServer(healthCheckHandler *HealthCheckHandler) *Server {
	engine := gin.Default()

	// 헬스 체크 엔드포인트 등록
	engine.GET("/health", healthCheckHandler.check)

	return &Server{
		engine: engine,
	}
}

func (s *Server) Run(addr string) error {
	return s.engine.Run(addr)
}

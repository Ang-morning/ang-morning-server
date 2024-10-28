package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthCheckHandler struct {
	// 필요한 의존성을 추가할 수 있습니다.
}

func NewHandler() *HealthCheckHandler {
	return &HealthCheckHandler{}
}

func (h *HealthCheckHandler) check(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

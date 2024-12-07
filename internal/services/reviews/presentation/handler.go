package presentation

import (
	httpCode "angmorning.com/internal/libs/http/http-code"
	httpError "angmorning.com/internal/libs/http/http-error"
	httpResponse "angmorning.com/internal/libs/http/http-response"
	"angmorning.com/internal/services/reviews/application"
	"angmorning.com/internal/services/reviews/command"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ReviewHandler struct {
	reviewService *application.ReviewService
}

func New(reviewService *application.ReviewService) *ReviewHandler {
	return &ReviewHandler{
		reviewService: reviewService,
	}
}

func (it *ReviewHandler) Router(r *gin.RouterGroup) {
	r.POST("/", it.Write)
}

func (it *ReviewHandler) Write(c *gin.Context) {
	var cmd command.WriteCommand
	if err := c.ShouldBindQuery(&cmd); err != nil {
		c.Error(httpError.New(httpCode.BadRequest, err.Error(), ""))
		return
	}
	// TODO:
	it.reviewService.Write(uuid.MustParse("328992c5-aef2-4aed-be41-c7d7a973e0fe"), cmd)

	c.JSON(httpCode.Created.Code, httpResponse.Response{Data: nil})
}

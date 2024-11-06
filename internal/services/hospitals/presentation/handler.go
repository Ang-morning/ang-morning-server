package presentation

import (
	httpCode "angmorning.com/internal/libs/http/http-code"
	httpError "angmorning.com/internal/libs/http/http-error"
	httpResponse "angmorning.com/internal/libs/http/http-response"
	"angmorning.com/internal/services/hospitals/application"
	"angmorning.com/internal/services/hospitals/command"
	"github.com/gin-gonic/gin"
)

type HospitalHandler struct {
	hospitalService *application.HospitalService
}

func New(hospitalService *application.HospitalService) *HospitalHandler {
	return &HospitalHandler{
		hospitalService: hospitalService,
	}
}

func (it *HospitalHandler) Router(r *gin.RouterGroup) {
	r.GET("/", it.List)

}

func (it *HospitalHandler) List(c *gin.Context) {
	var cmd command.ListCommand
	if err := c.ShouldBindQuery(&cmd); err != nil {
		c.Error(httpError.New(httpCode.BadRequest, err.Error(), ""))
		return
	}

	result, err := it.hospitalService.List(cmd)
	if err != nil {
		c.Error(httpError.New(httpCode.InternalServerError, err.Error(), ""))
		return
	}

	c.JSON(httpCode.Ok.Code, httpResponse.Response{Data: result})
}

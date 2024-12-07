package presentation

import (
	httpCode "angmorning.com/internal/libs/http/http-code"
	httpError "angmorning.com/internal/libs/http/http-error"
	httpResponse "angmorning.com/internal/libs/http/http-response"
	"angmorning.com/internal/services/users/application"
	"angmorning.com/internal/services/users/command"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *application.UserService
}

func New(userService *application.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (it *UserHandler) Router(r *gin.RouterGroup) {
	r.POST("/", it.oAuth)

}

func (it *UserHandler) oAuth(c *gin.Context) {
	var command command.OauthCommand
	if err := c.BindJSON(&command); err != nil {
		c.Error(httpError.New(httpCode.BadRequest, err.Error(), ""))
	}
	clientInfo := c.Request.Header.Get("User-Agent")

	res, err := it.userService.OAuth(command, clientInfo)
	if err != nil {
		c.Error(httpError.Wrap(err))
		return
	}

	c.JSON(httpCode.Created.Code, httpResponse.Response{Data: res.AccessToken})
}

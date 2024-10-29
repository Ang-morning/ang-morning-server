package presentation

import (
	"fmt"

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
		fmt.Println(err)
	}

	it.userService.OAuth(command)
}

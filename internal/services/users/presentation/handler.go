package presentation

import (
	"angmorning.com/internal/services/users/application"
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
	it.userService.OAuth()
}

package middlewares

import (
	"fmt"

	httpError "angmorning.com/internal/libs/http/http-error"
	httpResponse "angmorning.com/internal/libs/http/http-response"
	"github.com/gin-gonic/gin"
)

func ErrorHandler(ctx *gin.Context) {
	ctx.Next()
	err := ctx.Errors.Last()
	if err != nil {
		e := httpError.UnWrap(err.Err)

		//TODO: log error with something (e.g. Sentry, ELK, File, etc.)
		fmt.Println(e.Stack)

		ctx.JSON(e.Code, httpResponse.Response{Data: e.ClientMessage})
		return
	}
}

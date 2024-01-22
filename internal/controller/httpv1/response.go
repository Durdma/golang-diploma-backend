package httpv1

import "github.com/gin-gonic/gin"

type response struct {
	Message string `json:"message"`
}

func newResponse(ctx *gin.Context, statusCode int, msg string) {
	ctx.AbortWithStatusJSON(statusCode, response{msg})
}

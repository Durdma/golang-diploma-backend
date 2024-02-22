package httpv1

import (
	"github.com/gin-gonic/gin"
	"sas/pkg/logger"
)

type errorResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(ctx *gin.Context, statusCode int, message string) {
	logger.Error(message)
	ctx.AbortWithStatusJSON(statusCode, errorResponse{message})
}

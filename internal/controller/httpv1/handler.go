package httpv1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	error2 "sas/internal/models/error"
	"sas/internal/service"
)

type Handler struct {
	services *service.AdminsService
}

func NewHandler(services *service.AdminsService) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) Init() *gin.Engine {
	router := gin.Default()

	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK,
			error2.CustomError{
				Status: http.StatusOK,
				Msg:    "OK!",
			})
	})

	//router.GET("/sign-in")

	return router
}

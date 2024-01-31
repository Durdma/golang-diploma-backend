package controller

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	_ "sas/docs"
	"sas/internal/controller/httpv1"
	"sas/internal/service"
)

type Handler struct {
	universitiesService service.Universities
	editorsService      service.Editors
}

func NewHandler(universitiesService service.Universities, editorsService service.Editors) *Handler {
	return &Handler{
		universitiesService: universitiesService,
		editorsService:      editorsService,
	}
}

func (h *Handler) Init() *gin.Engine {
	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	h.initAPI(router)

	return router
}

func (h *Handler) initAPI(router *gin.Engine) {
	handlerV1 := httpv1.NewHandler(h.universitiesService, h.editorsService)
	api := router.Group("/api")
	{
		handlerV1.Init(api)
	}
}

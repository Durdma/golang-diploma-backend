package controller

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	_ "sas/docs"
	"sas/internal/controller/httpv1"
	"sas/internal/service"
	"sas/pkg/auth"
)

// Handler - Структура обработчика событий, главного
type Handler struct {
	universitiesService service.Universities
	editorsService      service.Editors
	tokenManager        auth.TokenManager
}

// NewHandler - Создание новой сущности обработчика событий
func NewHandler(universitiesService service.Universities, editorsService service.Editors, tokenManager auth.TokenManager) *Handler {
	return &Handler{
		universitiesService: universitiesService,
		editorsService:      editorsService,
		tokenManager:        tokenManager,
	}
}

// Init - Инициализация обработчика событий, добавление delevelopers роутов
func (h *Handler) Init() *gin.Engine {
	router := gin.Default() // Инициализируем стандартный маршрутизатор
	// Добавление нужных middleware
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	// Для отображения документации api
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Проверка работы api
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Добавляем все имеющиеся группы роутеров
	h.initAPI(router)

	return router
}

// initAPI - Объединение в более общую группу роутеров
func (h *Handler) initAPI(router *gin.Engine) {
	handlerV1 := httpv1.NewHandler(h.universitiesService, h.editorsService, h.tokenManager)
	api := router.Group("/api")
	{
		handlerV1.Init(api)
	}
}

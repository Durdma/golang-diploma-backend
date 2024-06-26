package controller

import (
	"context"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"sas/docs"
	_ "sas/docs"
	"sas/internal/controller/httpv1"
	"sas/internal/service"
	"sas/pkg/auth"
	"sas/pkg/logger"
)

// Handler - Структура обработчика событий, главного
type Handler struct {
	universitiesService service.Universities
	editorsService      service.Editors
	adminsService       service.Admins
	tokenManager        auth.TokenManager
	domainsService      service.Domains
	usersService        service.Users
	newsService         service.News
	docsService         service.Docs
}

// NewHandler - Создание новой сущности обработчика событий
func NewHandler(universitiesService service.Universities, editorsService service.Editors,
	adminsService service.Admins, tokenManager auth.TokenManager, domainsService service.Domains,
	usersService service.Users, newsService service.News, docsService service.Docs) *Handler {
	return &Handler{
		universitiesService: universitiesService,
		editorsService:      editorsService,
		adminsService:       adminsService,
		tokenManager:        tokenManager,
		domainsService:      domainsService,
		usersService:        usersService,
		newsService:         newsService,
		docsService:         docsService,
	}
}

// Init - Инициализация обработчика событий, добавление delevelopers роутов
func (h *Handler) Init(host string, port string) *gin.Engine {
	router := gin.Default() // Инициализируем стандартный маршрутизатор
	// Добавление нужных middleware
	router.Use(
		gin.Recovery(),
		gin.Logger(),
		cors.New(cors.Config{
			AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
			AllowCredentials: true,
			AllowHeaders:     []string{"content-type"},
			AllowOriginFunc: func(origin string) bool {
				origins, err := h.domainsService.GetAllDomains(context.Background())
				fmt.Println(origins)
				if err != nil {
					logger.Error("error while fetching allow origins from DB")
					return false
				}

				for _, orig := range origins {
					if "http://platform.localhost:3000" == origin {
						fmt.Println("===========================================")
						fmt.Println(orig)
						fmt.Println("===========================================")
						return true
					}

					if "http://"+orig.HTTPDomainName+".localhost:5000" == origin {
						fmt.Println("===========================================")
						fmt.Println(orig)
						fmt.Println("===========================================")
						return true
					}

				}

				return false
			},
		}),
	)

	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", host, port)

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
	handlerV1 := httpv1.NewHandler(h.universitiesService, h.editorsService, h.adminsService, h.tokenManager,
		h.domainsService, h.usersService, h.newsService, h.docsService)
	api := router.Group("/api")
	{
		handlerV1.Init(api)
	}
}

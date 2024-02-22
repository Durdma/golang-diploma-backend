package httpv1

import (
	"github.com/gin-gonic/gin"
	"sas/internal/service"
	"sas/pkg/auth"
)

// Handler - Структура обработчика событий, обработки поступающих запросов для сервисов
type Handler struct {
	universitiesService service.Universities // Сервис для работы с логикой университетов
	editorsService      service.Editors      // Сервис для работы с логикой редакторов
	tokenManager        auth.TokenManager
}

// NewHandler - Создание обработчика событий. На вход передаем уже инициализированные сервисы
func NewHandler(universitiesService service.Universities, editorsService service.Editors, tokenManager auth.TokenManager) *Handler {
	return &Handler{
		universitiesService: universitiesService,
		editorsService:      editorsService,
		tokenManager:        tokenManager,
	}
}

// Init - Инициализация обработчика событий. Подключаем все имеющиеся группы роутеров
func (h *Handler) Init(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	{
		h.initEditorsRoutes(v1)
	}
}

package httpv1

import (
	"github.com/gin-gonic/gin"
)

// initEditorsRoutes - Инициализация группы роутера для редакторов сайта университета
func (h *Handler) initEditorsRoutes(api *gin.RouterGroup) {
	// Группирует все маршруты редакторов
	editors := api.Group("/editors", h.setDomainFromRequest)
	{

		authenticated := editors.Group("/")
		{
			authenticated.GET("/news")
			authenticated.GET("/news/:id")
		}
	}
}

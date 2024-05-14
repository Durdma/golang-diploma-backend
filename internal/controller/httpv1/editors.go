package httpv1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sas/internal/models"
)

// initEditorsRoutes - Инициализация группы роутера для редакторов сайта университета
func (h *Handler) initEditorsRoutes(api *gin.RouterGroup) {
	// Группирует все маршруты редакторов
	editors := api.Group("/editors", h.setUniversityFromRequest)
	{

		authenticated := editors.Group("/", h.userIdentity)
		{
			authenticated.GET("/news", h.editorGetAllNews)
			authenticated.GET("/news/:id", h.editorsGetNewsById)
		}
	}
}

// @Summary Editor GetByHTTPName All News
// @Tags editors
// @Description editor get all news
// @ID editorGetAllNews
// @Accept json
// @Produce json
// @Success 200 {array} models.News
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /editors/news [get]
func (h *Handler) editorGetAllNews(ctx *gin.Context) {
	univ, err := getUniversityFromContext(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	news := make([]models.News, 0)
	for _, n := range univ.News {
		if n.Published {
			news = append(news, n)
		}
	}

	ctx.JSON(http.StatusOK, news)
}

// @Summary Editor GetByHTTPName News By ID
// @Tags editors
// @Description editor get news by id
// @ID editorsGetNewsById
// @Accept json
// @Produce json
// @Param id path string true "news id"
// @Success 200 {object} models.News
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /editors/news/{id} [get]
func (h *Handler) editorsGetNewsById(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid id param")
		return
	}

	univ, err := getUniversityFromContext(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	for _, n := range univ.News {
		if n.Published && n.ID.Hex() == id {
			ctx.JSON(http.StatusOK, n)
		}
	}

	newErrorResponse(ctx, http.StatusBadRequest, "not found")
}

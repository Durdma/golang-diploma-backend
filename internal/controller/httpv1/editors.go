package httpv1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sas/internal/service"
	"sas/pkg/logger"
)

// initEditorsRoutes - Инициализация группы роутера для редакторов сайта университета
func (h *Handler) initEditorsRoutes(api *gin.RouterGroup) {
	// Группирует все маршруты редакторов
	editors := api.Group("/editors", h.setUniversityFromRequest())
	{
		editors.POST("/sign-up", h.editorsSignUp)    // Регистрация нового редактора сайта
		editors.POST("/sign-in")                     // Вход редактора на сайт
		editors.GET("/verify/:hash", h.editorVerify) // Подтверждение учетной записи редактора ПОКА ИЗМЕНЕНО НА GET (POST изначально)
	}
}

// editorsSignUpInput - структура, в которую парсится тело запроса на регистрацию редактора
type editorsSignUpInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"hash"`
}

// editorsSignUp - Парсит тело полученного запроса в структуру и получает домен университета, от которого пришел запрос
func (h *Handler) editorsSignUp(ctx *gin.Context) {
	var input editorsSignUpInput
	if err := ctx.BindJSON(&input); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	univ, err := getUniversityFromContext(ctx)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if err := h.editorsService.SignUp(ctx.Request.Context(), service.EditorSignUpInput{
		Name:         input.Name,
		Email:        input.Email,
		Password:     input.Password,
		UniversityID: univ.ID,
	}); err != nil {
		logger.Error(err)

		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusCreated)
}

// editorVerify - Подтверждение создания учетной записи редактора
func (h *Handler) editorVerify(ctx *gin.Context) {
	hash := ctx.Param("hash")
	if hash == "" {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := h.editorsService.Verify(ctx.Request.Context(), hash); err != nil {
		logger.Error(err)
		ctx.AbortWithStatus(http.StatusBadRequest)
	}

	ctx.Status(http.StatusOK)
}

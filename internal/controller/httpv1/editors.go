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
		editors.POST("/sign-in", h.editorSignIn)     // Вход редактора на сайт
		editors.GET("/verify/:hash", h.editorVerify) // Подтверждение учетной записи редактора ПОКА ИЗМЕНЕНО НА GET (POST изначально)
	}
}

// editorsSignUpInput - структура, в которую парсится тело запроса на регистрацию редактора
type editorsSignUpInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"hash" binding:"required"`
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

type editorsSignInInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type tokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (h *Handler) editorSignIn(ctx *gin.Context) {
	var inp editorsSignInInput
	if err := ctx.BindJSON(&inp); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	univ, err := getUniversityFromContext(ctx)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	res, err := h.editorsService.SignIn(ctx.Request.Context(), service.EditorSignInInput{
		UniversityID: univ.ID,
		Email:        inp.Email,
		Password:     inp.Password,
	})
	if err != nil {
		logger.Error(err)
	}

	ctx.JSON(http.StatusOK, tokenResponse{
		AccessToken:  res.AccessToken,
		RefreshToken: res.RefreshToken,
	})
}

type refreshInput struct {
	Token string `json:"token" binding:"required"`
}

func (h *Handler) editorRefresh(ctx *gin.Context) {
	var inp refreshInput
	if err := ctx.BindJSON(&inp); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	univ, err := getUniversityFromContext(ctx)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	res, err := h.editorsService.RefreshTokens(ctx.Request.Context(), univ.ID, inp.Token)
	if err != nil {
		logger.Error(err)

		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, tokenResponse{
		AccessToken:  res.AccessToken,
		RefreshToken: res.RefreshToken,
	})
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

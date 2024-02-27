package httpv1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sas/internal/models"
	"sas/internal/service"
	"sas/pkg/logger"
)

// initEditorsRoutes - Инициализация группы роутера для редакторов сайта университета
func (h *Handler) initEditorsRoutes(api *gin.RouterGroup) {
	// Группирует все маршруты редакторов
	editors := api.Group("/editors", h.setUniversityFromRequest)
	{
		editors.POST("/sign-up", h.editorsSignUp) // Регистрация нового редактора сайта
		editors.POST("/sign-in", h.editorSignIn)  // Вход редактора на сайт
		editors.POST("/auth/refresh", h.editorRefresh)
		editors.GET("/verify/:hash", h.editorVerify) // Подтверждение учетной записи редактора ПОКА ИЗМЕНЕНО НА GET (POST изначально)

		authenticated := editors.Group("/", h.userIdentity)
		{
			authenticated.GET("/news", h.editorGetAllNews)
			authenticated.GET("/news/:id", h.editorsGetNewsById)
		}
	}
}

// editorsSignUpInput - структура, в которую парсится тело запроса на регистрацию редактора
type editorsSignUpInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"hash" binding:"required"`
}

// @Summary Editor SignUp
// @Tags editors
// @Description create editor account
// @ID editorSignUp
// @Accept json
// @Produce json
// @Param input body editorsSignUpInput true "sign up info"
// @Success 201 {string} string "ok"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /editors/sign-up [post]
// editorsSignUp - Парсит тело полученного запроса в структуру и получает домен университета, от которого пришел запрос
func (h *Handler) editorsSignUp(ctx *gin.Context) {
	var input editorsSignUpInput
	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid input body")
		return
	}

	univ, err := getUniversityFromContext(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	if err := h.editorsService.SignUp(ctx.Request.Context(), service.EditorSignUpInput{
		Name:         input.Name,
		Email:        input.Email,
		Password:     input.Password,
		UniversityID: univ.ID,
	}); err != nil {
		logger.Error(err)

		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
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

// @Summary Editor SignIn
// @Tags editors
// @Description editor sign in
// @ID editorSignIn
// @Accept json
// @Produce json
// @Param input body editorsSignInInput true "sign in info"
// @Success 200 {object} tokenResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /editors/sign-in [post]
func (h *Handler) editorSignIn(ctx *gin.Context) {
	var inp editorsSignInInput
	if err := ctx.BindJSON(&inp); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid input body")
		return
	}

	univ, err := getUniversityFromContext(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	res, err := h.editorsService.SignIn(ctx.Request.Context(), service.EditorSignInInput{
		UniversityID: univ.ID,
		Email:        inp.Email,
		Password:     inp.Password,
	})
	if err != nil {
		logger.Error(err)

		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, tokenResponse{
		AccessToken:  res.AccessToken,
		RefreshToken: res.RefreshToken,
	})
}

type refreshInput struct {
	Token string `json:"token" binding:"required"`
}

// @Summary Editor Refresh Token
// @Security EditorsAuth
// @Tags editors
// @Description editor refresh tokens
// @ID editorRefresh
// @Accept json
// @Produce json
// @Param input body refreshInput true "sign up info"
// @Success 200 {object} tokenResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /editors/refresh [post]
func (h *Handler) editorRefresh(ctx *gin.Context) {
	var inp refreshInput
	if err := ctx.BindJSON(&inp); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid input body")
		return
	}

	univ, err := getUniversityFromContext(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	res, err := h.editorsService.RefreshTokens(ctx.Request.Context(), univ.ID, inp.Token)
	if err != nil {
		logger.Error(err)

		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, tokenResponse{
		AccessToken:  res.AccessToken,
		RefreshToken: res.RefreshToken,
	})
}

// @Summary Editor Verify Registration
// @Tags editors
// @Description editor verify registration
// @ID editorVerify
// @Accept json
// @Produce json
// @Param code path string true "verification code"
// @Success 200 {object} tokenResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /editors/verify/{code} [post]
// editorVerify - Подтверждение создания учетной записи редактора
func (h *Handler) editorVerify(ctx *gin.Context) {
	hash := ctx.Param("hash")
	if hash == "" {
		newErrorResponse(ctx, http.StatusBadRequest, "code is empty")
		return
	}

	if err := h.editorsService.Verify(ctx.Request.Context(), hash); err != nil {
		logger.Error(err)
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.Status(http.StatusOK)
}

// @Summary Editor Get All News
// @Tags editors
// @Security EditorsAuth
// @Description editor get all news
// @ID editorGetAllNews
// @Accept json
// @Produce json
// @Success 200 {array} university.News
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

// @Summary Editor Get News By ID
// @Tags editors
// @Security EditorsAuth
// @Description editor get news by id
// @ID editorsGetNewsById
// @Accept json
// @Produce json
// @Param id path string true "news id"
// @Success 200 {object} university.News
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

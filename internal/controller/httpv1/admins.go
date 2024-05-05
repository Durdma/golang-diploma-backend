package httpv1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sas/internal/renderer"
	"sas/internal/service"
	"sas/internal/view"
	"sas/pkg/logger"
)

func (h *Handler) initAdminsRoutes(api *gin.RouterGroup) {
	admins := api.Group("/admins", h.setDomainFromRequest)
	{
		admins.POST("/sign-in", h.adminSignIn)
		admins.GET("/sign-in", h.getAdminSignIn)
		admins.POST("/auth/refresh", h.adminRefresh)
		admins.GET("/verify/:hash", h.adminVerify)
		admins.GET("/logout", h.logOut)
	}

	//authenticated := admins.Group("/", h.userIdentity)
	authenticated := admins.Group("/")
	{
		domains := authenticated.Group("/sites")
		{
			domains.GET("")

			domains.GET("/new")
			domains.POST("/new", h.postDomain)

			domains.GET("/:id")
			domains.PATCH("/:id")
			domains.DELETE("/:id")
		}

		adminsGroup := authenticated.Group("/admins")
		{
			adminsGroup.POST("/new", h.adminsSignUp)
			admins.GET("/new")
		}

		// Не в приоритете
		requestsGroup := authenticated.Group("/requests")
		{
			requestsGroup.GET("")

			requestsGroup.GET("/new")
			requestsGroup.POST("/new")

			requestsGroup.GET("/:id")
			requestsGroup.PATCH("/:id")
			requestsGroup.DELETE("/:id")
		}

		employeesGroup := authenticated.Group("/employees")
		{
			employeesGroup.GET("")

			employeesGroup.GET("/new")
			employeesGroup.POST("/new")

			employeesGroup.GET("/:id")
			employeesGroup.PATCH("/:id")
			employeesGroup.DELETE("/:id")
		}

		// Не в приоритете
		notificationsGroup := authenticated.Group("/notifications")
		{
			notificationsGroup.GET("")

			notificationsGroup.GET("/new")
			notificationsGroup.POST("/new")

			notificationsGroup.GET("/:id")
		}
	}
}

type adminsSignUpInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) adminsSignUp(ctx *gin.Context) {
	var input adminsSignUpInput
	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid input body")
		return
	}

	//univ, err := getUniversityFromContext(ctx)
	//if err != nil {
	//	newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	//	return
	//}

	// TODO implement error if cant find user in db
	if err := h.adminsService.SignUp(ctx.Request.Context(), service.AdminSignUpInput{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	}); err != nil {
		logger.Error(err)

		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.Status(http.StatusOK)
}

type adminsSignInInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// TODO implement resolve domain
func (h *Handler) getAdminSignIn(ctx *gin.Context) {
	resp := renderer.New(ctx.Request.Context(), http.StatusOK, view.Login())
	ctx.Render(http.StatusOK, resp)
}

func (h *Handler) adminSignIn(ctx *gin.Context) {
	domain, ex := ctx.Get("db_domain")
	if !ex {
		newErrorResponse(ctx, http.StatusBadRequest, "not exists")
		return
	}

	// TODO refactor for platform subdomain
	if domain != "test1" {
		newErrorResponse(ctx, http.StatusForbidden, "forbidden")
		return
	}

	var input adminsSignInInput
	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid body input")
		return
	}

	res, err := h.adminsService.SignIn(ctx.Request.Context(), service.AdminSignInInput{
		Email:    input.Email,
		Password: input.Password,
	})
	if err != nil {
		logger.Error(err)

		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.SetCookie("access_token", res.AccessToken, res.AccessTokenTTL, "/", "localhost", false, true)
	ctx.Status(http.StatusOK)
}

func (h *Handler) adminRefresh(ctx *gin.Context) {
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

	res, err := h.adminsService.RefreshTokens(ctx.Request.Context(), univ.ID, inp.Token)
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

func (h *Handler) adminVerify(ctx *gin.Context) {
	hash := ctx.Param("hash")
	if hash == "" {
		newErrorResponse(ctx, http.StatusBadRequest, "code is empty")
		return
	}

	if err := h.adminsService.Verify(ctx.Request.Context(), hash); err != nil {
		logger.Error(err)
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *Handler) logOut(ctx *gin.Context) {
	ctx.SetCookie("access_token", "", 0, "/", "localhost", false, true)
	ctx.Status(http.StatusOK)
}

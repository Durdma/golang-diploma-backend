package httpv1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sas/internal/service"
	"sas/pkg/logger"
)

func (h *Handler) initAdminsRoutes(api *gin.RouterGroup) {
	admins := api.Group("/admins")
	{
		admins.POST("/sign-in", h.adminSignIn)
		admins.GET("/sign-in")
		admins.POST("/auth/refresh", h.adminRefresh)
		admins.GET("/verify/:hash", h.adminVerify)
	}

	//authenticated := admins.Group("/", h.userIdentity)
	authenticated := admins.Group("/")
	{
		sitesGroup := authenticated.Group("/sites")
		{
			sitesGroup.GET("")

			sitesGroup.GET("/new")
			sitesGroup.POST("/new")

			sitesGroup.GET("/:id")
			sitesGroup.PATCH("/:id")
			sitesGroup.DELETE("/:id")
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

func (h *Handler) adminSignIn(ctx *gin.Context) {
	var input adminsSignInInput
	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid body input")
		return
	}

	//univ, err := getUniversityFromContext(ctx)
	//if err != nil {
	//	newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	//	return
	//}

	res, err := h.adminsService.SignIn(ctx.Request.Context(), service.AdminSignInInput{
		Email:    input.Email,
		Password: input.Password,
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

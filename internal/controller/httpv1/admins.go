package httpv1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sas/internal/service"
	"sas/pkg/logger"
)

func (h *Handler) initAdminsRoutes(api *gin.RouterGroup) {
	admins := api.Group("/admins", h.setUniversityFromRequest)
	{
		admins.POST("/sign-up")
		admins.POST("/sign-in")
		admins.POST("/auth/refresh")
		admins.GET("/verify/:hash")
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

	univ, err := getUniversityFromContext(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

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

	univ, err := getUniversityFromContext(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
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

	ctx.JSON(http.StatusOK, tokenResponse{
		AccessToken:  res.AccessToken,
		RefreshToken: res.RefreshToken,
	})
}

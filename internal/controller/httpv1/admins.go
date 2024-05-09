package httpv1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sas/internal/service"
	"sas/pkg/logger"
)

func (h *Handler) initAdminsRoutes(api *gin.RouterGroup) {
	admins := api.Group("/admins", h.setDomainFromRequest)
	{

	}

	//authenticated := admins.Group("/", h.userIdentity)
	authenticated := admins.Group("/", h.setUserFromRequest)
	{
		domains := authenticated.Group("/sites")
		{
			domains.GET("", h.getAllSites)

			domains.GET("/new", h.getNewSite)
			domains.POST("/new", h.postDomain) //TODO refactor to POST site

			domains.GET("/:id", h.getSite)
			domains.PATCH("/:id", h.patchSite)
			domains.DELETE("/:id", h.deleteSite)
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
			employeesGroup.GET("", h.getAllEditors)

			employeesGroup.GET("/new")
			employeesGroup.POST("/new")

			employeesGroup.GET("/:id")
			employeesGroup.PATCH("/:id", h.patchEditor)
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

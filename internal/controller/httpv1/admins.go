package httpv1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
			domains.GET("", h.getAllDomains)

			domains.POST("", h.postSite)

			domains.GET("/:id", h.getSite)
			domains.PATCH("/:id", h.patchSite)
			domains.DELETE("/:id", h.deleteSite)
		}

		adminsGroup := authenticated.Group("/admins")
		{
			adminsGroup.POST("", h.adminsSignUp)
			adminsGroup.GET("", h.getAllAdmins)
			adminsGroup.PATCH("/:id", h.patchAdmin)
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
			employeesGroup.POST("", h.postNewEditor)

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
	if code, err := getAdminsPermissions(ctx); err != nil {
		newErrorResponse(ctx, code, err.Error())
		return
	}

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

	fmt.Println("OK")
	ctx.Status(http.StatusOK)
}

func (h *Handler) getAllAdmins(ctx *gin.Context) {
	if code, err := getAdminsPermissions(ctx); err != nil {
		newErrorResponse(ctx, code, err.Error())
		return
	}

	getQueryToContext(ctx)

	admins, err := h.adminsService.GetAll(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	fmt.Println(admins)
	ctx.JSON(http.StatusOK, admins)
}

func getQueryToContext(ctx *gin.Context) {
	domainName := ctx.Query("domain_name")
	if domainName != "" {
		ctx.Set("domain_name", domainName)
	}

	name := ctx.Query("name")
	if name != "" {
		ctx.Set("name", name)
	}

	university := ctx.Query("university")
	if university != "" {
		fmt.Println(university)
		ctx.Set("university", university)
	}

	sort := ctx.Query("sort")
	if sort != "" {
		ctx.Set("sort", sort)
	}

	verify := ctx.Query("verify")
	if verify != "" {
		ctx.Set("verify", verify)
	}

	block := ctx.Query("block")
	if block != "" {
		ctx.Set("block", block)
	}

	visible := ctx.Query("visible")
	if visible != "" {
		ctx.Set("visible", visible)
	}
}

type patchAdminInput struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Verify   bool   `json:"verify"`
	Block    bool   `json:"block"`
}

func (h *Handler) patchAdmin(ctx *gin.Context) {
	if code, err := getAdminsPermissions(ctx); err != nil {
		newErrorResponse(ctx, code, err.Error())
		return
	}

	userId := ctx.Param("id")

	verify, verifyEx := ctx.GetQuery("verify")
	block, blockEx := ctx.GetQuery("block")

	switch {
	case verifyEx && blockEx:
		newErrorResponse(ctx, http.StatusBadRequest, "error query params")
		return
	case blockEx && !verifyEx:
		err := h.adminsService.ChangeAdminBlockStatus(ctx, userId, block)
		if err != nil {
			newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
			return
		}
	case verifyEx && !blockEx:
		err := h.adminsService.ChangeAdminVerifyStatus(ctx, userId, verify)
		if err != nil {
			newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
			return
		}
	case !verifyEx && !blockEx:
		var userInput patchAdminInput
		if err := ctx.BindJSON(&userInput); err != nil {
			newErrorResponse(ctx, http.StatusBadRequest, err.Error())
			return
		}

		domainIdCtx, _ := ctx.Get("db_domain")
		domainId := domainIdCtx.(primitive.ObjectID)

		err := h.adminsService.UpdateAdmin(ctx, service.UpdateAdminInput{
			Id:       userInput.Id,
			Name:     userInput.Name,
			Email:    userInput.Email,
			Password: userInput.Password,
			DomainId: domainId.Hex(),
			Verify:   userInput.Verify,
			Block:    userInput.Block,
		})
		if err != nil {
			newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
			return
		}
	default:
		newErrorResponse(ctx, http.StatusInternalServerError, "something wrong")
		return
	}

	ctx.Status(http.StatusOK)
}

package httpv1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sas/internal/service"
	"sas/pkg/logger"
)

func (h *Handler) initEditorsRoutes(api *gin.RouterGroup) {
	editors := api.Group("/editors", h.setUniversityFromRequest())
	{
		editors.POST("/sign-up", h.editorsSignUp)
		editors.POST("/sign-in")
		editors.POST("/verify/:hash")
	}
}

type editorsSignUpInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

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
	}

	ctx.Status(http.StatusCreated)
}

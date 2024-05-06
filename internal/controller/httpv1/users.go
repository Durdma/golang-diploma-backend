package httpv1

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"sas/internal/service"
)

func (h *Handler) initUsersRoutes(api *gin.RouterGroup) {
	user := api.Group("/auth", h.setDomainFromRequest)
	{
		user.POST("/sign-in", h.signIn)
		user.GET("/sign-in", h.getSignIn)
		user.GET("/refresh", h.refresh)
		user.GET("/verify/:hash", h.verify)
		user.GET("/logout", h.logOut)
	}
}

type userSignInInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signIn(ctx *gin.Context) {
	domain, ex := ctx.Get("db_domain")
	if !ex {
		newErrorResponse(ctx, http.StatusBadRequest, "no db_domain")
		return
	}

	var input userSignInInput
	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid body input")
		return
	}

	res, err := h.usersService.SignIn(ctx, service.SignInInput{
		Email:    input.Email,
		Password: input.Password,
		Domain:   domain.(primitive.ObjectID),
	})
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.SetCookie("access_token", res.AccessToken, res.AccessTokenTTL, "/", "localhost", false, true)
	ctx.SetCookie("refresh_token", res.RefreshToken, res.RefreshTokenTTL, "/", "localhost", false, true)
	ctx.Status(http.StatusOK)
}

func (h *Handler) getSignIn(ctx *gin.Context) {

}

type refreshInput struct {
	Token string `json:"token" binding:"required"`
}

func (h *Handler) refresh(ctx *gin.Context) {
	domain, ex := ctx.Get("db_domain")
	if !ex {
		newErrorResponse(ctx, http.StatusBadRequest, "no db_domain")
		return
	}

	refreshToken, err := ctx.Request.Cookie("refresh_token")
	if err != nil {
		newErrorResponse(ctx, http.StatusForbidden, "no refresh_token")
		return
	}

	res, err := h.usersService.RefreshTokens(ctx, domain.(primitive.ObjectID), refreshToken.Value)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.SetCookie("access_token", res.AccessToken, res.AccessTokenTTL, "/", "localhost", false, true)
	ctx.SetCookie("refresh_token", res.RefreshToken, res.RefreshTokenTTL, "/", "localhost", false, true)
	ctx.Status(http.StatusOK)
}

func (h *Handler) verify(ctx *gin.Context) {
	domain, ex := ctx.Get("db_domain")
	if !ex {
		newErrorResponse(ctx, http.StatusBadRequest, "no db_domain")
		return
	}

	hash := ctx.Param("hash")
	if hash == "" {
		newErrorResponse(ctx, http.StatusBadRequest, "code is empty")
		return
	}

	if err := h.usersService.Verify(ctx, domain.(primitive.ObjectID), hash); err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *Handler) logOut(ctx *gin.Context) {
	ctx.SetCookie("access_token", "", 0, "/", "localhost", false, true)
	ctx.SetCookie("refresh_token", "", 0, "/", "localhost", false, true)
	ctx.Status(http.StatusOK)
}

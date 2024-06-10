package httpv1

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"sas/pkg/logger"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
	universityCtx       = "university"
)

func (h *Handler) setDomainFromRequest(ctx *gin.Context) {
	hostName := strings.Split(ctx.Request.Header.Get("Origin"), "://")
	if len(hostName) != 2 {
		logger.Error(errors.New("error length of Origin not equal to 2"))
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	domains := strings.Split(hostName[1], ".")
	if len(domains) != 2 {
		logger.Error(errors.New("error length of domains not equal to 2"))
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	subDomain := domains[0]

	resp, err := h.domainsService.GetByHTTPName(ctx, subDomain)
	if err != nil {
		logger.Error(err.Error())
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	ctx.Set("db_domain", resp.ID)
	ctx.Set("dom", resp.HTTPDomainName)
	ctx.Set("university_id", resp.SiteId)
}

func (h *Handler) setUserFromRequest(ctx *gin.Context) {
	accessToken, err := ctx.Cookie("access_token")
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	userId, err := h.tokenManager.Parse(accessToken)
	if err != nil {
		newErrorResponse(ctx, http.StatusForbidden, "no such user")
		return
	}

	user, err := h.usersService.GetUserById(ctx, userId)
	if err != nil {
		newErrorResponse(ctx, http.StatusForbidden, "no such user")
		return
	}

	ctx.Set("user_id", user.ID)
	ctx.Set("is_admin", user.IsAdmin)
	ctx.Set("verified", user.Verification.Verified)
	ctx.Set("domain_id", user.DomainId)
}

func getAdminsPermissions(ctx *gin.Context) (int, error) {
	domain, ex := ctx.Get("dom")
	if !ex {
		return http.StatusUnauthorized, errors.New("no dom ctx")
	}

	if domain.(string) != "platform" {
		return http.StatusForbidden, errors.New("incorrect domain")
	}

	isAdmin, ex := ctx.Get("is_admin")
	if !ex {
		return http.StatusUnauthorized, errors.New("no is_admin ctx")
	}

	if !isAdmin.(bool) {
		return http.StatusForbidden, errors.New("access forbidden")
	}

	verified, ex := ctx.Get("verified")
	if !ex {
		return http.StatusUnauthorized, errors.New("access forbidden")
	}

	if !verified.(bool) {
		return http.StatusForbidden, errors.New("access forbidden")
	}

	return 0, nil
}

func getEditorsPermissions(ctx *gin.Context) (int, error) {
	origin, ex := ctx.Get("db_domain")
	if !ex {
		return http.StatusUnauthorized, errors.New("no dom ctx")
	}

	domainId, ex := ctx.Get("domain_id")
	if !ex {
		return http.StatusUnauthorized, errors.New("no domain_id")
	}

	if origin.(primitive.ObjectID) != domainId.(primitive.ObjectID) {
		return http.StatusForbidden, errors.New("access forbidden")
	}

	verified, ex := ctx.Get("verified")
	if !ex {
		return http.StatusUnauthorized, errors.New("access forbidden")
	}

	if !verified.(bool) {
		return http.StatusForbidden, errors.New("access forbidden")
	}

	return 0, nil
}

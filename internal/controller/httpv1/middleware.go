package httpv1

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"sas/internal/models"
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
}

// setUniversityFromRequest - Получение домена, с которого пришел запрос и обращение к нужному университету
func (h *Handler) setUniversityFromRequest(ctx *gin.Context) {
	domainName := strings.Split(ctx.Request.Host, ":")[0]

	univ, err := h.universitiesService.GetByDomain(ctx.Request.Context(), domainName)
	if err != nil {
		logger.Error(err)
		ctx.AbortWithStatus(http.StatusForbidden)
		return
	}

	ctx.Set(universityCtx, univ)
}

// getUniversityFromContext - Получение имени университета из контекста запроса
func getUniversityFromContext(ctx *gin.Context) (models.University, error) {
	value, ex := ctx.Get(universityCtx)
	if !ex {
		return models.University{}, errors.New("university is missing from context")
	}

	univ, ok := value.(models.University)
	if !ok {
		return models.University{}, errors.New("failed to convert value from ctx to university.University")
	}

	return univ, nil
}

func (h *Handler) userIdentity(ctx *gin.Context) {
	header := ctx.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponse(ctx, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		newErrorResponse(ctx, http.StatusUnauthorized, "invalid auth header")
		return
	}

	if len(headerParts[1]) == 0 {
		newErrorResponse(ctx, http.StatusUnauthorized, "token is empty")
		return
	}

	userId, err := h.tokenManager.Parse(headerParts[1])
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	ctx.Set(userCtx, userId)
}

func getAdminsPermissions(ctx *gin.Context) (int, error) {
	domain, ex := ctx.Get("dom")
	if !ex {
		return http.StatusUnauthorized, errors.New("no dom ctx")
	}

	if domain.(string) != "test1" {
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

func getUserId(ctx *gin.Context) (string, error) {
	id, ok := ctx.Get(userCtx)
	if !ok {
		return "", errors.New("user id not found")
	}

	idStr, ok := id.(string)
	if !ok {
		return "", errors.New("user id is of invalid type")
	}

	return idStr, nil
}

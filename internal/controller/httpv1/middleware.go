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

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
	universityCtx = "university"
)

// setUniversityFromRequest - Получение домена, с которого пришел запрос и обращение к нужному университету
func (h *Handler) setUniversityFromRequest() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		domainName := strings.Split(ctx.Request.Host, ":")[0]
		//fmt.Println(domainName)
		univ, err := h.universitiesService.GetByDomain(ctx.Request.Context(), domainName)
		if err != nil {
			logger.Error(err)

			ctx.AbortWithStatus(http.StatusForbidden)
			return
		}

		ctx.Set(universityCtx, univ)
	}
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

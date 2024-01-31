package httpv1

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sas/internal/models/university"
	"sas/pkg/logger"
	"strings"
)

const (
	universityCtx = "university"
)

func (h *Handler) setUniversityFromRequest() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		domainName := strings.Split(ctx.Request.Host, ":")[0]
		fmt.Println(domainName)
		univ, err := h.universitiesService.GetByDomain(ctx.Request.Context(), domainName)
		if err != nil {
			logger.Error(err)

			ctx.AbortWithStatus(http.StatusForbidden)
			return
		}

		ctx.Set(universityCtx, univ)
	}
}

func getUniversityFromContext(ctx *gin.Context) (university.University, error) {
	value, ex := ctx.Get(universityCtx)
	if !ex {
		return university.University{}, errors.New("university is missing from context")
	}

	univ, ok := value.(university.University)
	if !ok {
		return university.University{}, errors.New("failed to convert value from ctx to university.University")
	}

	return univ, nil
}

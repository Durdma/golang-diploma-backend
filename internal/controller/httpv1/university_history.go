package httpv1

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"sas/internal/models"
)

func (h *Handler) getUniversityHistory(ctx *gin.Context) {
	universityId, ex := ctx.Get("university_id")
	if !ex {
		newErrorResponse(ctx, http.StatusForbidden, "access forbidden")
		return
	}

	university, err := h.universitiesService.GetByUniversityId(ctx, universityId.(primitive.ObjectID))
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, university.History)
}

func (h *Handler) postUniversityHistory(ctx *gin.Context) {
	universityId, ex := ctx.Get("university_id")
	if !ex {
		newErrorResponse(ctx, http.StatusForbidden, "access forbidden")
		return
	}

	var history models.History
	if err := ctx.BindJSON(&history); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	err := h.universitiesService.SetUniversityHistory(ctx, universityId.(primitive.ObjectID), history)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *Handler) patchUniversityHistory(ctx *gin.Context) {

}

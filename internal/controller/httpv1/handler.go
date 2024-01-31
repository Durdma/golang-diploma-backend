package httpv1

import (
	"github.com/gin-gonic/gin"
	"sas/internal/service"
)

type Handler struct {
	universitiesService service.Universities
	editorsService      service.Editors
}

func NewHandler(universitiesService service.Universities, editorsService service.Editors) *Handler {
	return &Handler{
		universitiesService: universitiesService,
		editorsService:      editorsService,
	}
}

func (h *Handler) Init(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	{
		h.initEditorsRoutes(v1)
	}
}

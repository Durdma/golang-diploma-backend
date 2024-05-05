package httpv1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sas/internal/service"
	"sas/pkg/logger"
)

//func (h *Handler) initAdminsDomainsRouter(api *gin.RouterGroup) {
//	domains := api.Group("/admins", h.setDomainFromRequest)
//	{
//		authenticated := domains.Group("/domains", h.userIdentity)
//		{
//			authenticated.GET("")
//			authenticated.GET("/new")
//			authenticated.POST("/new")
//			authenticated.DELETE("/:id")
//		}
//	}
//}

type POSTDomainInput struct {
	HTTPDomain string `json:"http_domain" binding:"required"`
}

func (h *Handler) postDomain(ctx *gin.Context) {
	var input POSTDomainInput
	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid body input")
		return
	}

	if err := h.domainsService.AddDomain(ctx, service.DomainInput{
		HTTPDomain: input.HTTPDomain,
	}); err != nil {
		logger.Error(err)

		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.Status(http.StatusOK)
}

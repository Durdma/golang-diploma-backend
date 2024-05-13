package httpv1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sas/internal/service"
	"sas/pkg/logger"
)

type POSTDomainInput struct {
	HTTPDomain string `json:"http_domain" binding:"required"`
}

// TODO refactor to POST site
func (h *Handler) postDomain(ctx *gin.Context) {
	if code, err := getAdminsPermissions(ctx); err != nil {
		newErrorResponse(ctx, code, err.Error())
		return
	}

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

func (h *Handler) getNewSite(ctx *gin.Context) {

}

type POSTSiteInput struct {
	Name           string `json:"name"`
	ShortName      string `json:"short_name"`
	HTTPDomainName string `json:"http_domain_name"`
}

func (h *Handler) postNewSite(ctx *gin.Context) {
	if code, err := getAdminsPermissions(ctx); err != nil {
		newErrorResponse(ctx, code, err.Error())
		return
	}

	var input POSTSiteInput
	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid body input")
		return
	}

	universityId, err := h.sitesService.AddNewSite(ctx, service.SiteInput{
		Name:           input.Name,
		ShortName:      input.ShortName,
		HTTPDomainName: input.HTTPDomainName,
	})
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	if err := h.domainsService.AddDomain(ctx, service.DomainInput{
		HTTPDomain: input.HTTPDomainName,
		SiteId:     universityId,
		Name:       input.Name,
		ShortName:  input.ShortName,
		Verified:   false,
	}); err != nil {
		logger.Error(err)

		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *Handler) getAllDomains(ctx *gin.Context) {
	domains, err := h.domainsService.GetAllDomains(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, domains)
}

func (h *Handler) getSite(ctx *gin.Context) {

}

func (h *Handler) patchSite(ctx *gin.Context) {

}

func (h *Handler) deleteSite(ctx *gin.Context) {

}

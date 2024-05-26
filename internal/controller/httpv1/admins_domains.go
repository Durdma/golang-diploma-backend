package httpv1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sas/internal/service"
)

type postSiteInput struct {
	DomainName string `json:"domain_name"`
	ShortName  string `json:"short_name"`
	HTTPName   string `json:"http_name"`
	Verify     bool   `json:"verify"`
	Visible    bool   `json:"visible"`
}

func (h *Handler) postSite(ctx *gin.Context) {
	if code, err := getAdminsPermissions(ctx); err != nil {
		newErrorResponse(ctx, code, err.Error())
		return
	}

	var input postSiteInput
	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid body input")
		return
	}

	domainId, err := h.domainsService.AddDomain(ctx, service.DomainInput{
		DomainName: input.DomainName,
		ShortName:  input.ShortName,
		HTTPName:   input.HTTPName,
		Verify:     input.Verify,
		Visible:    input.Visible,
	})
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	universityId, err := h.universitiesService.AddUniversity(ctx, domainId, input.DomainName, input.ShortName)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.domainsService.AddUniversityId(ctx, domainId, universityId)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *Handler) getNewSite(ctx *gin.Context) {

}

func (h *Handler) getAllDomains(ctx *gin.Context) {
	getQueryToContext(ctx)

	domains, err := h.domainsService.GetAllDomains(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, domains)
}

func (h *Handler) getSite(ctx *gin.Context) {

}

type patchDomainInput struct {
	Id         string `json:"id"`
	DomainName string `json:"domain_name"`
	ShortName  string `json:"short_name"`
	Verified   bool   `json:"verified"`
	Visible    bool   `json:"visible"`
}

func (h *Handler) patchSite(ctx *gin.Context) {
	if code, err := getAdminsPermissions(ctx); err != nil {
		newErrorResponse(ctx, code, err.Error())
		return
	}

	domainId := ctx.Param("id")

	verify, verifyEx := ctx.GetQuery("verify")
	visible, visibleEx := ctx.GetQuery("visible")

	switch {
	case verifyEx && visibleEx:
		newErrorResponse(ctx, http.StatusBadRequest, "error query params")
		return
	case verifyEx && !visibleEx:
		err := h.domainsService.ChangeSiteVerifyStatus(ctx, domainId, verify)
		if err != nil {
			newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
			return
		}
	case visibleEx && !verifyEx:
		err := h.domainsService.ChangeSiteVisibleStatus(ctx, domainId, visible)
		if err != nil {
			newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
			return
		}
	case !verifyEx && !visibleEx:
		var domainInput patchDomainInput
		if err := ctx.BindJSON(&domainInput); err != nil {
			newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
			return
		}

		err := h.domainsService.UpdateDomain(ctx, service.UpdateDomainInput{
			Id:         domainInput.Id,
			DomainName: domainInput.DomainName,
			ShortName:  domainInput.ShortName,
			Visible:    domainInput.Visible,
			Verified:   domainInput.Verified,
		})
		if err != nil {
			newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
			return
		}
	default:
		newErrorResponse(ctx, http.StatusInternalServerError, "something wrong")
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *Handler) deleteSite(ctx *gin.Context) {

}

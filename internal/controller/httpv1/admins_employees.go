package httpv1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// TODO add to models.Editor field univ name
func (h *Handler) getAllEditors(ctx *gin.Context) {
	//domain, ex := ctx.GetByHTTPName("db_domain")
	//if !ex {
	//	newErrorResponse(ctx, http.StatusBadRequest, "no db_domain")
	//	return
	//
	// TODO refactor verification to function
	dom, ex := ctx.Get("dom")
	if !ex {
		newErrorResponse(ctx, http.StatusBadRequest, "no dom")
		return
	}

	if dom.(string) != "test1" {
		newErrorResponse(ctx, http.StatusForbidden, "no permissions")
		return
	}

	isAdmin, ex := ctx.Get("is_admin")
	if !ex {
		newErrorResponse(ctx, http.StatusForbidden, "no is_admin")
		return
	}

	verified, ex := ctx.Get("verified")
	if !ex {
		newErrorResponse(ctx, http.StatusForbidden, "no is_admin")
		return
	}

	if !isAdmin.(bool) || !verified.(bool) {
		newErrorResponse(ctx, http.StatusForbidden, "no permissions")
		return
	}

	editors, err := h.usersService.GetAllEditors(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	for idx, editor := range editors {
		domain, err := h.domainsService.GetById(ctx, editor.DomainId)
		if err != nil {
			newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
			return
		}
		fmt.Println(domain)
		editors[idx].DomainName = domain.DomainName
	}

	ctx.JSON(http.StatusOK, editors)
}

func (h *Handler) patchEditor(ctx *gin.Context) {
	dom, ex := ctx.Get("dom")
	if !ex {
		newErrorResponse(ctx, http.StatusBadRequest, "no dom")
		return
	}

	if dom.(string) != "test1" {
		newErrorResponse(ctx, http.StatusForbidden, "no permissions")
		return
	}

	isAdmin, ex := ctx.Get("is_admin")
	if !ex {
		newErrorResponse(ctx, http.StatusBadRequest, "no is_admin")
		return
	}

	verified, ex := ctx.Get("verified")
	if !ex {
		newErrorResponse(ctx, http.StatusBadRequest, "no is_admin")
		return
	}

	if !isAdmin.(bool) || !verified.(bool) {
		newErrorResponse(ctx, http.StatusForbidden, "no permissions")
		return
	}

	userId := ctx.Param("id")

	verify, verifyEx := ctx.GetQuery("verify")
	block, blockEx := ctx.GetQuery("block")

	switch {
	case verifyEx && blockEx:
		newErrorResponse(ctx, http.StatusBadRequest, "error query params")
		return
	case blockEx:
		err := h.editorsService.ChangeEditorBlockStatus(ctx, userId, block)
		if err != nil {
			newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
			return
		}
	case verifyEx:
		err := h.editorsService.ChangeEditorVerifyStatus(ctx, userId, verify)
		if err != nil {
			newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
			return
		}
	default:

	}

	ctx.Status(http.StatusOK)
}

func (h *Handler) postNewEditor(ctx *gin.Context) {

}

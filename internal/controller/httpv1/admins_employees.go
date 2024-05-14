package httpv1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"sas/internal/service"
)

func (h *Handler) getAllEditors(ctx *gin.Context) {
	if code, err := getAdminsPermissions(ctx); err != nil {
		newErrorResponse(ctx, code, err.Error())
		return
	}

	name := ctx.Query("name")
	if name != "" {
		ctx.Set("name", name)
	}

	university := ctx.Query("university")
	if university != "" {
		ctx.Set("university", university)
	}

	sort := ctx.Query("sort")
	if sort != "" {
		ctx.Set("sort", sort)
	}

	verify := ctx.Query("verify")
	if verify != "" {
		ctx.Set("verify", verify)
	}

	block := ctx.Query("block")
	if block != "" {
		ctx.Set("block", block)
	}

	editors, err := h.editorsService.GetAllEditors(ctx)
	if err != nil {
		fmt.Println("from here 1")
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	for idx, editor := range editors {
		domain, err := h.domainsService.GetById(ctx, editor.DomainId)
		if err != nil {
			fmt.Println("from here 2")
			newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
			return
		}
		fmt.Println(domain)
		editors[idx].DomainName = domain.DomainName
	}

	fmt.Println(editors)
	ctx.JSON(http.StatusOK, editors)
}

type patchEditorInput struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	DomainName string `json:"domain_name"`
	DomainId   string `json:"domain_id"`
	Verify     bool   `json:"verify"`
	Block      bool   `json:"block"`
}

func (h *Handler) patchEditor(ctx *gin.Context) {
	if code, err := getAdminsPermissions(ctx); err != nil {
		newErrorResponse(ctx, code, err.Error())
		return
	}

	userId := ctx.Param("id")

	verify, verifyEx := ctx.GetQuery("verify")
	block, blockEx := ctx.GetQuery("block")

	switch {
	case verifyEx && blockEx:
		newErrorResponse(ctx, http.StatusBadRequest, "error query params")
		return
	case blockEx && !verifyEx:
		err := h.editorsService.ChangeEditorBlockStatus(ctx, userId, block)
		if err != nil {
			newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
			return
		}
	case verifyEx && !blockEx:
		err := h.editorsService.ChangeEditorVerifyStatus(ctx, userId, verify)
		if err != nil {
			newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
			return
		}
	case !verifyEx && !blockEx:
		var userInput patchEditorInput
		if err := ctx.BindJSON(&userInput); err != nil {
			newErrorResponse(ctx, http.StatusBadRequest, err.Error())
			return
		}

		fmt.Println(ctx.Request.Body)

		err := h.editorsService.UpdateEditor(ctx, service.UpdateEditorInput{
			Id:         userInput.Id,
			Name:       userInput.Name,
			Email:      userInput.Email,
			Password:   userInput.Password,
			DomainName: userInput.DomainName,
			DomainId:   userInput.DomainId,
			Verify:     userInput.Verify,
			Block:      userInput.Block,
		})
		if err != nil {
			newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
			return
		}
	default:
		newErrorResponse(ctx, http.StatusBadRequest, "something wrong")
		return
	}

	ctx.Status(http.StatusOK)
}

type postEditorInput struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	DomainName string `json:"domain_name"`
	DomainId   string `json:"domain_id"`
	Verify     bool   `json:"verify"`
	Block      bool   `json:"block"`
}

func (h *Handler) postNewEditor(ctx *gin.Context) {
	fmt.Println(ctx.Request.Body)
	if code, err := getAdminsPermissions(ctx); err != nil {
		newErrorResponse(ctx, code, err.Error())
		return
	}

	var editor postEditorInput
	if err := ctx.BindJSON(&editor); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	domainId, err := primitive.ObjectIDFromHex(editor.DomainId)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	domain, err := h.domainsService.GetById(ctx, domainId)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	if err := h.editorsService.SignUp(ctx, service.EditorSignUpInput{
		Name:       editor.Name,
		Email:      editor.Email,
		Password:   editor.Password,
		DomainName: domain.DomainName,
		DomainId:   domain.ID,
		Verify:     editor.Verify,
		Block:      editor.Block,
	}); err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	}

	ctx.Status(http.StatusOK)
}

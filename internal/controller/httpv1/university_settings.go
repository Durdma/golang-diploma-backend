package httpv1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"sas/internal/models"
)

func (h *Handler) getLogoImage(ctx *gin.Context) {
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

	if university.Settings.HeaderImage != "" {
		ctx.File(university.Settings.Label)
	}

	ctx.Status(http.StatusNotFound)
}

func (h *Handler) getHeaderImage(ctx *gin.Context) {
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

	//TODO refactor like newsheaderImage
	if university.Settings.HeaderImage != "" {
		ctx.File(university.Settings.HeaderImage)
	}

	ctx.Status(http.StatusNotFound)
}

type getUniversityResponse struct {
	University models.University `json:"university"`
	Domain     models.Domain     `json:"domain"`
}

func (h *Handler) getUniversity(ctx *gin.Context) {
	if code, err := getEditorsPermissions(ctx); err != nil {
		newErrorResponse(ctx, code, err.Error())
		return
	}

	universityId, ex := ctx.Get("university_id")
	if !ex {
		newErrorResponse(ctx, http.StatusInternalServerError, "something wrong")
		return
	}

	university, err := h.universitiesService.GetByUniversityId(ctx, universityId.(primitive.ObjectID))
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	domain, err := h.domainsService.GetById(ctx, university.DomainId)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	resp := getUniversityResponse{
		University: university,
		Domain:     domain,
	}

	ctx.JSON(http.StatusOK, resp)

}

func (h *Handler) getColors(ctx *gin.Context) {
	universityId, ex := ctx.Get("university_id")
	if !ex {
		newErrorResponse(ctx, http.StatusInternalServerError, "something wrong")
		return
	}

	colors, err := h.universitiesService.GetUniversityColors(ctx, universityId.(primitive.ObjectID))
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	fmt.Println(")()())()())()()()()()()()()()()()()()()()()()()()()()()()()()()()()()()()(")
	fmt.Println(colors)
	fmt.Println(")()())()())()()()()()()()()()()()()()()()()()()()()()()()()()()()()()()()(")
	ctx.JSON(http.StatusOK, colors)
}

func (h *Handler) getCSS(ctx *gin.Context) {
	universityId, ex := ctx.Get("university_id")
	if !ex {
		newErrorResponse(ctx, http.StatusInternalServerError, "something wrong")
		return
	}

	ctx.File("../../static/css/" + universityId.(primitive.ObjectID).Hex() + "_" + "css" + ".css")
}

type patchCSSInput struct {
	MainColor                string `json:"main_color"`
	MainColorHover           string `json:"main_color_hover"`
	MainFooterFontColor      string `json:"main_footer_font_color"`
	MainFooterFontColorHover string `json:"main_footer_font_color_hover"`
	MainFooterBgColor        string `json:"main_footer_bg_color"`
}

func (h *Handler) patchCSS(ctx *gin.Context) {
	universityId, ex := ctx.Get("university_id")
	if !ex {
		fmt.Println("here1")
		newErrorResponse(ctx, http.StatusInternalServerError, "something wrong")
		return
	}

	var input patchCSSInput
	if err := ctx.BindJSON(&input); err != nil {
		fmt.Println("here2")
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	colors := map[string]string{
		"main_color":                   input.MainColor,
		"main_color_hover":             input.MainColorHover,
		"main_footer_font_color":       input.MainFooterFontColor,
		"main_footer_font_color_hover": input.MainFooterFontColorHover,
		"main_footer_bg_color":         input.MainFooterBgColor,
	}

	err := h.universitiesService.PatchUniversityCSS(ctx, universityId.(primitive.ObjectID), colors)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.Status(http.StatusOK)
}

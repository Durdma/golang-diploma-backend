package httpv1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"sas/internal/service"
	"strings"
)

func (h *Handler) getAllNews(ctx *gin.Context) {
	universityId, ex := ctx.Get("university_id")
	if !ex {
		newErrorResponse(ctx, http.StatusForbidden, "access forbidden")
		return
	}

	news, err := h.newsService.GetAllNews(ctx, universityId.(primitive.ObjectID))
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, news)
}

type postNewsInput struct {
	Header      string `json:"header"`
	Description string `json:"description"`
	Body        string `json:"body"`
	CreatedBy   string `json:"created_by"`
}

func (h *Handler) postNews(ctx *gin.Context) {
	universityId, ex := ctx.Get("university_id")
	if !ex {
		newErrorResponse(ctx, http.StatusForbidden, "access forbidden")
		return
	}

	var newsInput postNewsInput
	if err := ctx.BindJSON(&newsInput); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	recordId, err := h.newsService.AddNews(ctx, service.NewsInput{
		Header:       newsInput.Header,
		Description:  newsInput.Description,
		Body:         newsInput.Body,
		CreatedBy:    newsInput.CreatedBy,
		UniversityId: universityId.(primitive.ObjectID),
	})
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"record_id": recordId,
	})
}

func (h *Handler) getImage(ctx *gin.Context) {
	_, ex := ctx.Get("university_id")
	if !ex {
		newErrorResponse(ctx, http.StatusForbidden, "access forbidden")
		return
	}

	imageURL := ctx.Param("name")
	if imageURL == "" {
		newErrorResponse(ctx, http.StatusBadRequest, "no imageURL")
		return
	}

	ctx.File("../../static/media/images/news_header_image/" + imageURL)
}

func (h *Handler) setImage(ctx *gin.Context) {
	universityId, ex := ctx.Get("university_id")
	if !ex {
		newErrorResponse(ctx, http.StatusForbidden, "access forbidden")
		return
	}

	recordId := ctx.Param("name")
	if recordId == "" {
		newErrorResponse(ctx, http.StatusBadRequest, "no imageURL")
		return
	}

	file, err := ctx.FormFile("image")
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	fmt.Println(file.Header)
	fmt.Println(file.Filename)

	imageType := strings.Split(file.Filename, ".")[1]

	file.Filename = universityId.(primitive.ObjectID).Hex() + "_" + recordId + "_" + "newsHeaderImage" + "." + imageType

	err = ctx.SaveUploadedFile(file, "..\\..\\static\\media\\images\\news_header_image\\"+file.Filename)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.newsService.AddHeaderImageURL(ctx, recordId, file.Filename)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.Status(http.StatusOK)
}

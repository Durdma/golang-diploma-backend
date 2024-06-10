package httpv1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"sas/internal/service"
	"strings"
	"time"
)

type postDocsInput struct {
	Header          string `json:"header"`
	Description     string `json:"description"`
	PublicationDate string `json:"publication_date"`
	CreatedBy       string `json:"created_by"`
}

func (h *Handler) postDocs(ctx *gin.Context) {
	universityId, ex := ctx.Get("university_id")
	if !ex {
		newErrorResponse(ctx, http.StatusForbidden, "access forbidden")
		return
	}

	var input postDocsInput
	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	timestmp, err := time.Parse(time.DateOnly, input.PublicationDate) //2006-01-02
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	docId, err := h.docsService.AddDocs(ctx, service.DocsInput{
		UniversityId:    universityId.(primitive.ObjectID),
		Header:          input.Header,
		Description:     input.Description,
		PublicationDate: timestmp,
		CreatedBy:       input.CreatedBy,
	})
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"doc_id": docId,
	})
}

type postStudyPlan struct {
	Header          string `json:"header"`
	Description     string `json:"description"`
	Code            string `json:"code"`
	Magistrate      bool   `json:"magistrate"`
	Enrollee        bool   `json:"enrollee"`
	PublicationDate string `json:"publication_date"`
	CreatedBy       string `json:"created_by"`
}

func (h *Handler) postStudyPlanDocs(ctx *gin.Context) {
	universityId, ex := ctx.Get("university_id")
	if !ex {
		fmt.Println("here 1")
		newErrorResponse(ctx, http.StatusForbidden, "access forbidden")
		return
	}

	var input postStudyPlan
	if err := ctx.BindJSON(&input); err != nil {
		fmt.Println("here 2")
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	timestmp, err := time.Parse(time.DateOnly, input.PublicationDate) //2006-01-02
	if err != nil {
		fmt.Println("here 3")
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	docId, err := h.docsService.AddDocs(ctx, service.DocsInput{
		UniversityId:    universityId.(primitive.ObjectID),
		Header:          input.Header,
		Description:     input.Description,
		Code:            input.Code,
		Magistrate:      input.Magistrate,
		Enrollee:        input.Enrollee,
		PublicationDate: timestmp,
		CreatedBy:       input.CreatedBy,
	})
	if err != nil {
		fmt.Println("here 4")
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"doc_id": docId,
	})
}

func (h *Handler) setDocs(ctx *gin.Context) {
	universityId, ex := ctx.Get("university_id")
	if !ex {
		newErrorResponse(ctx, http.StatusForbidden, "access forbidden")
		return
	}

	docId := ctx.Param("doc_id")
	if docId == "" {
		newErrorResponse(ctx, http.StatusBadRequest, "no doc_id")
		return
	}

	file, err := ctx.FormFile("doc")
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	docType := strings.Split(file.Filename, ".")[1]

	file.Filename = universityId.(primitive.ObjectID).Hex() + "_" + docId + "." + docType

	err = ctx.SaveUploadedFile(file, "..\\..\\static\\media\\docs\\university_docs\\"+file.Filename)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.docsService.AddDocsURL(ctx, docId, file.Filename)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *Handler) getAllUniversityDocs(ctx *gin.Context) {
	universityId, ex := ctx.Get("university_id")
	if !ex {
		newErrorResponse(ctx, http.StatusForbidden, "access forbidden")
		return
	}

	docs, err := h.docsService.GetAllUniversityDocs(ctx, universityId.(primitive.ObjectID))
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	fmt.Println(docs)

	ctx.JSON(http.StatusOK, docs)
}

func (h *Handler) getDocs(ctx *gin.Context) {
	_, ex := ctx.Get("university_id")
	if !ex {
		newErrorResponse(ctx, http.StatusForbidden, "access forbidden")
		return
	}

	docsURL := ctx.Param("doc_id")
	if docsURL == "" {
		newErrorResponse(ctx, http.StatusBadRequest, "no docsURL")
		return
	}

	filePath := "../../static/media/docs/university_docs/" + docsURL
	fileName := docsURL // Имя файла будет таким же, как и в docsURL

	ctx.FileAttachment(filePath, fileName)
}

func (h *Handler) getAllBachelors(ctx *gin.Context) {
	universityId, ex := ctx.Get("university_id")
	if !ex {
		newErrorResponse(ctx, http.StatusForbidden, "access forbidden")
		return
	}

	docs, err := h.docsService.GetAllBachelors(ctx, universityId.(primitive.ObjectID))
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	fmt.Println(docs)

	ctx.JSON(http.StatusOK, docs)
}

func (h *Handler) getAllMags(ctx *gin.Context) {
	universityId, ex := ctx.Get("university_id")
	if !ex {
		newErrorResponse(ctx, http.StatusForbidden, "access forbidden")
		return
	}

	docs, err := h.docsService.GetAllMags(ctx, universityId.(primitive.ObjectID))
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, docs)
}

func (h *Handler) getAllEnrollsDocs(ctx *gin.Context) {
	universityId, ex := ctx.Get("university_id")
	if !ex {
		newErrorResponse(ctx, http.StatusForbidden, "access forbidden")
		return
	}

	docs, err := h.docsService.GetAllEnrollsDocs(ctx, universityId.(primitive.ObjectID))
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, docs)
}

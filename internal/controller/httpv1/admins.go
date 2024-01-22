package httpv1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sas/internal/models/admin"
	"sas/internal/service"
)

type AdminsController struct {
	service *service.AdminsService
}

func NewAdminsController(service *service.AdminsService) *AdminsController {
	return &AdminsController{service: service}
}

func (ac *AdminsController) SignUpAdmin(ctx *gin.Context) {
	var data admin.Admin

	if err := ctx.BindJSON(&data); err != nil {
		newResponse(ctx, http.StatusBadRequest, "Cant unmarshall JSON!")

		return
	}

	//domain resolve middleware

}

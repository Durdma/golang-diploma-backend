package service

import (
	"sas/internal/models"
	"sas/internal/repository"
)

type AdminsService struct {
	repository *repository.TmpRepo
}

func NewAdminService(repo *repository.TmpRepo) *AdminsService {
	return &AdminsService{repository: repo}
}

func (as *AdminsService) SignUp(data *models.AdminRecord) (string, *models.CustomError) {
	//Some validation

	return as.repository.AddAdmin(data)
}

func (as *AdminsService) SignIn(email string, username string, password string) {

}

func (as *AdminsService) GetAllAdmins() ([]*models.AdminRecord, *models.CustomError) {
	return as.repository.GetAdmins()
}

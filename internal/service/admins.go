package service

import (
	"sas/internal/models/admin"
	error2 "sas/internal/models/error"
	"sas/internal/repository"
)

type AdminsService struct {
	repository *repository.TmpRepo
}

func NewAdminService(repo *repository.TmpRepo) *AdminsService {
	return &AdminsService{repository: repo}
}

func (as *AdminsService) SignUp(data *admin.Admin) (string, *error2.CustomError) {
	//Some validation

	return as.repository.AddAdmin(data)
}

func (as *AdminsService) SignIn(email string, username string, password string) {

}

func (as *AdminsService) GetAllAdmins() ([]*admin.Admin, *error2.CustomError) {
	return as.repository.GetAdmins()
}

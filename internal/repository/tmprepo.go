package repository

import (
	"net/http"
	"sas/internal/models"
)

type TmpRepo struct {
	repo []*models.AdminRecord
}

func NewTmpRepo() *TmpRepo {
	return &TmpRepo{repo: make([]*models.AdminRecord, 0)}
}

func (r *TmpRepo) AddAdmin(data *models.AdminRecord) (string, *models.CustomError) {
	for _, rec := range r.repo {
		if rec.Email == data.Email {
			return "", &models.CustomError{
				Status: http.StatusBadRequest,
				Msg:    "You can't register more than one account on email.",
			}
		}

		if rec.Username == data.Username {
			return "", &models.CustomError{
				Status: http.StatusBadRequest,
				Msg:    "Username is already used!",
			}
		}
	}

	data.ID = len(r.repo) + 1

	r.repo = append(r.repo, data)

	return "Added!", nil
}

func (r *TmpRepo) GetAdmins() ([]*models.AdminRecord, *models.CustomError) {
	if len(r.repo) == 0 {
		return nil, &models.CustomError{
			Status: http.StatusBadRequest,
			Msg:    "Empty DB!",
		}
	}
	return r.repo, nil
}

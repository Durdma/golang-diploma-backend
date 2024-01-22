package repository

import (
	"net/http"
	"sas/internal/models/admin"
	error2 "sas/internal/models/error"
)

type TmpRepo struct {
	repo []*admin.Admin
}

func NewTmpRepo() *TmpRepo {
	return &TmpRepo{repo: make([]*admin.Admin, 0)}
}

func (r *TmpRepo) AddAdmin(data *admin.Admin) (string, *error2.CustomError) {
	for _, rec := range r.repo {
		if rec.Email == data.Email {
			return "", &error2.CustomError{
				Status: http.StatusBadRequest,
				Msg:    "You can't register more than one account on email.",
			}
		}

		if rec.Username == data.Username {
			return "", &error2.CustomError{
				Status: http.StatusBadRequest,
				Msg:    "Username is already used!",
			}
		}
	}

	data.ID = len(r.repo) + 1

	r.repo = append(r.repo, data)

	return "Added!", nil
}

func (r *TmpRepo) GetAdmins() ([]*admin.Admin, *error2.CustomError) {
	if len(r.repo) == 0 {
		return nil, &error2.CustomError{
			Status: http.StatusBadRequest,
			Msg:    "Empty DB!",
		}
	}
	return r.repo, nil
}

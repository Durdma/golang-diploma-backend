package service

import (
	"context"
	"sas/internal/models/university"
	"sas/internal/repository"
)

type EditorsService struct {
	repo repository.Editors
}

func NewEditorsService(repo repository.Editors) *EditorsService {
	return &EditorsService{
		repo: repo,
	}
}

func (s *EditorsService) SignIn(ctx context.Context, email string, password string) (string, error) {
	return "", nil
}

func (s *EditorsService) SignUp(ctx context.Context, input EditorSignUpInput) error {
	editor := university.Editor{
		Name:         input.Name,
		Password:     input.Password,
		Email:        input.Email,
		UniversityID: input.UniversityID,
		Verification: university.Verification{
			Hash: "",
		},
	}

	return s.repo.Create(ctx, editor)
}

func (s *EditorsService) Verify(ctx context.Context, hash string) error {
	return nil
}

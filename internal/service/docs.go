package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"sas/internal/models"
	"sas/internal/repository"
	"time"
)

type DocsService struct {
	repo repository.Docs
}

func NewDocsService(repo repository.Docs) *DocsService {
	return &DocsService{repo: repo}
}

func (s *DocsService) AddDocs(ctx context.Context, input DocsInput) (primitive.ObjectID, error) {
	createdBy, err := primitive.ObjectIDFromHex(input.CreatedBy)
	if err != nil {
		return primitive.ObjectID{}, err
	}

	return s.repo.Create(ctx, models.Docs{
		UniversityId:    input.UniversityId,
		Header:          input.Header,
		Description:     input.Description,
		Code:            input.Code,
		Magistrate:      input.Magistrate,
		Enrollee:        input.Enrollee,
		DocURL:          "",
		PublicationDate: input.PublicationDate,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
		CreatedBy:       createdBy,
		UpdatedBy:       createdBy,
	})
}

func (s *DocsService) AddDocsURL(ctx context.Context, docIdString string, docURL string) error {
	docId, err := primitive.ObjectIDFromHex(docIdString)
	if err != nil {
		return err
	}

	return s.repo.AddDocsURL(ctx, docId, docURL)
}

func (s *DocsService) GetAllUniversityDocs(ctx context.Context, universityId primitive.ObjectID) ([]models.Docs, error) {
	return s.repo.GetAllUniversityDocs(ctx, universityId)
}

func (s *DocsService) GetAllBachelors(ctx context.Context, universityId primitive.ObjectID) ([]models.Docs, error) {
	return s.repo.GetAllBachelors(ctx, universityId)
}

func (s *DocsService) GetAllMags(ctx context.Context, universityId primitive.ObjectID) ([]models.Docs, error) {
	return s.repo.GetAllMags(ctx, universityId)
}

func (s *DocsService) GetAllEnrollsDocs(ctx context.Context, universityId primitive.ObjectID) ([]models.Docs, error) {
	return s.repo.GetAllEnrollsDocs(ctx, universityId)
}

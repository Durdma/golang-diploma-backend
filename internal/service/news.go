package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"sas/internal/models"
	"sas/internal/repository"
	"time"
)

type NewsService struct {
	repo repository.News
}

func NewNewsService(repo repository.News) *NewsService {
	return &NewsService{repo: repo}
}

func (s *NewsService) AddNews(ctx context.Context, input NewsInput) (primitive.ObjectID, error) {
	createdBy, err := primitive.ObjectIDFromHex(input.CreatedBy)
	if err != nil {
		return primitive.ObjectID{}, err
	}

	return s.repo.Create(ctx, models.News{
		DomainId:    input.UniversityId,
		Header:      input.Header,
		Description: input.Description,
		Body:        input.Body,
		ImageURL:    "",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Published:   true,
		CreatedBy:   createdBy,
		UpdatedBy:   createdBy,
	})
}

func (s *NewsService) GetAllNews(ctx context.Context, domainId primitive.ObjectID) ([]models.News, error) {
	return s.repo.GetAllNews(ctx, domainId)
}

func (s *NewsService) AddHeaderImageURL(ctx context.Context, recordId string, imageURL string) error {
	id, err := primitive.ObjectIDFromHex(recordId)
	if err != nil {
		return err
	}

	return s.repo.AddHeaderImageURL(ctx, id, imageURL)
}

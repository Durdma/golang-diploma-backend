package service

import (
	"context"
	"sas/internal/models"
	"sas/internal/repository"
	"sas/pkg/cache"
	"sas/pkg/logger"
)

// UniversitiesService - Структура для работы с сервисом университетов
type UniversitiesService struct {
	repo  repository.Universities
	cache cache.Cache
}

// NewUniversitiesService - Создание структуры сервиса университетов
func NewUniversitiesService(repo repository.Universities, cache cache.Cache) *UniversitiesService {
	return &UniversitiesService{
		repo:  repo,
		cache: cache,
	}
}

// GetByDomain - Получение из БД записи об университете по полученному домену
func (s *UniversitiesService) GetByDomain(ctx context.Context, domainName string) (models.University, error) {
	if value, err := s.cache.Get(domainName); err == nil {
		return value.(models.University), nil
	}

	logger.Info(domainName)

	univ, err := s.repo.GetByDomain(ctx, domainName)
	if err != nil {
		return models.University{}, err
	}

	s.cache.Set(domainName, univ)

	return univ, nil
}

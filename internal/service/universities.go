package service

import (
	"context"
	"sas/internal/models/university"
	"sas/internal/repository"
	"sas/pkg/cache"
	"sas/pkg/logger"
)

type UniversitiesService struct {
	repo  repository.Universities
	cache cache.Cache
}

func NewUniversitiesService(repo repository.Universities, cache cache.Cache) *UniversitiesService {
	return &UniversitiesService{
		repo:  repo,
		cache: cache,
	}
}

func (s *UniversitiesService) GetByDomain(ctx context.Context, domainName string) (university.University, error) {
	if value, err := s.cache.Get(domainName); err == nil {
		return value.(university.University), nil
	}

	logger.Info(domainName)

	univ, err := s.repo.GetByDomain(ctx, domainName)
	if err != nil {
		return university.University{}, err
	}

	s.cache.Set(domainName, univ)

	return univ, nil
}

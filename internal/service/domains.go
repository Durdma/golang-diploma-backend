package service

import (
	"context"
	"sas/internal/models"
	"sas/internal/repository"
	"strings"
)

type DomainsService struct {
	repo repository.Domains
}

func NewDomainsService(repo repository.Domains) *DomainsService {
	return &DomainsService{
		repo: repo,
	}
}

func (s *DomainsService) AddDomain(ctx context.Context, input DomainInput) error {
	var domain models.Domain

	domain.HTTPDomainName = input.HTTPDomain
	domain.DBDomainName = strings.ReplaceAll(input.HTTPDomain, "-", "_")
	domain.Deleted = false
	domain.Visible = true //TODO change on false later

	if err := s.repo.Create(ctx, domain); err != nil {
		return err
	}

	return nil
}

func (s *DomainsService) GetAllDomains(ctx context.Context) ([]models.Domain, error) {
	return s.repo.GetAllDomains(ctx)
}

func (s *DomainsService) GetDomain(ctx context.Context, input string) (models.Domain, error) {
	return s.repo.Get(ctx, input)
}

func (s *DomainsService) DeleteDomain(ctx context.Context, input string) error {
	return nil
}

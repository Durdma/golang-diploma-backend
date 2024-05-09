package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	return s.repo.Create(ctx, models.Domain{
		SiteId:         input.SiteId,
		HTTPDomainName: input.HTTPDomain,
		DBDomainName:   strings.ReplaceAll(input.HTTPDomain, "-", "_"),
		DomainName:     input.Name,
		ShortName:      input.ShortName,
		Visible:        true, //TODO change on false later
		Deleted:        false,
		Verified:       input.Verified,
	})
}

func (s *DomainsService) DeleteDomain(ctx context.Context, domainId primitive.ObjectID) error {
	return nil
}

func (s *DomainsService) GetByHTTPName(ctx context.Context, domain string) (models.Domain, error) {
	return s.repo.GetByHTTPName(ctx, domain)
}

func (s *DomainsService) GetById(ctx context.Context, domainId primitive.ObjectID) (models.Domain, error) {
	return s.repo.GetById(ctx, domainId)
}

func (s *DomainsService) GetAllDomains(ctx context.Context) ([]models.Domain, error) {
	return s.repo.GetAllDomains(ctx)
}

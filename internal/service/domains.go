package service

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"sas/internal/models"
	"sas/internal/repository"
	"strings"
	"time"
)

type DomainsService struct {
	repo repository.Domains
}

func NewDomainsService(repo repository.Domains) *DomainsService {
	return &DomainsService{
		repo: repo,
	}
}

func (s *DomainsService) AddDomain(ctx context.Context, input DomainInput) (primitive.ObjectID, error) {
	return s.repo.Create(ctx, models.Domain{
		HTTPDomainName: input.HTTPName,
		DBDomainName:   strings.ReplaceAll(input.HTTPName, "-", "_"),
		DomainName:     input.DomainName,
		ShortName:      input.ShortName,
		Visible:        input.Visible,
		Verified:       input.Verify,
		RegisteredAt:   time.Now(),
		LastUpdate:     time.Now(),
	})
}

func (s *DomainsService) ChangeSiteVisibleStatus(ctx context.Context, domainId string, state string) error {
	stateBool := false
	if state == "true" {
		stateBool = true
	} else {
		if state == "false" {
			stateBool = false
		} else {
			return errors.New("incorrect state")
		}
	}

	return s.repo.ChangeVisibleStatus(ctx, domainId, stateBool)
}

func (s *DomainsService) ChangeSiteVerifyStatus(ctx context.Context, domainId string, state string) error {
	stateBool := false
	if state == "true" {
		stateBool = true
	} else {
		if state == "false" {
			stateBool = false
		} else {
			return errors.New("incorrect state")
		}
	}

	return s.repo.ChangeVerificationStatus(ctx, domainId, stateBool)
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

func (s *DomainsService) GetByDomainName(ctx context.Context, domainName string) (models.Domain, error) {
	return s.repo.GetByDomainName(ctx, domainName)
}

func (s *DomainsService) GetAllDomains(ctx context.Context) ([]models.Domain, error) {
	return s.repo.GetAllDomains(ctx)
}

type UpdateDomainInput struct {
	Id         string
	DomainName string
	ShortName  string
	Visible    bool
	Verified   bool
}

func (s *DomainsService) UpdateDomain(ctx context.Context, newDomain UpdateDomainInput) error {
	domainId, err := primitive.ObjectIDFromHex(newDomain.Id)
	if err != nil {
		return err
	}

	oldDomain, err := s.repo.GetById(ctx, domainId)
	if err != nil {
		return err
	}

	oldDomain.DomainName = newDomain.DomainName
	oldDomain.ShortName = newDomain.ShortName
	oldDomain.Visible = newDomain.Visible
	oldDomain.Verified = newDomain.Verified

	err = s.repo.UpdateDomain(ctx, oldDomain)

	return err
}

func (s *DomainsService) AddUniversityId(ctx context.Context, domainId primitive.ObjectID, universityId primitive.ObjectID) error {
	return s.repo.AddUniversityId(ctx, domainId, universityId)
}

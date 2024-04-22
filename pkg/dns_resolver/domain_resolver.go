package dns_resolver

import (
	"context"
	"sas/internal/repository"
)

type DNS interface {
	AddDomain(ctx context.Context, domain string) error
	DelDomain(ctx context.Context, domain string) error
	GetDomain(ctx context.Context, domain string) (string, error)
}

type Resolver struct {
	repo repository.DNS
}

func NewResolver(db repository.DNS) *Resolver {
	return &Resolver{
		repo: db,
	}
}

func (r *Resolver) AddDomain(ctx context.Context, domain string) error {

	err := r.repo.Create(ctx, domain)

	return err
}

func (r *Resolver) DelDomain(ctx context.Context, domain string) error {

	err := r.repo.Delete(ctx, domain)

	return err
}

func (r *Resolver) GetDomain(ctx context.Context, domain string) (string, error) {

	domainName, err := r.repo.Get(ctx, domain)

	return domainName, err
}

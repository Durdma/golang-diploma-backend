package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"sas/internal/models/university"
	"sas/internal/repository"
	"sas/pkg/cache"
)

type Universities interface {
	GetByDomain(ctx context.Context, domainName string) (university.University, error)
}

// EditorSignUpInput  TODO Взято из примера для понимания, при добавлении редакторов переписать
type EditorSignUpInput struct {
	Name         string
	Email        string
	Password     string
	UniversityID primitive.ObjectID
}

type Editors interface {
	SignIn(ctx context.Context, email string, password string) (string, error)
	SignUp(ctx context.Context, input EditorSignUpInput) error
	Verify(ctx context.Context, hash string) error
}

type Services struct {
	Universities Universities
	Editors      Editors
}

func NewServices(repos *repository.Repositories, cache cache.Cache) *Services {
	return &Services{
		Universities: NewUniversitiesService(repos.Universities, cache),
		Editors:      NewEditorsService(repos.Editors),
	}
}

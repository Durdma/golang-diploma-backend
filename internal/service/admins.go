package service

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"sas/internal/models"
	"sas/internal/repository"
	"sas/pkg/auth"
	"sas/pkg/hash"
	"time"
)

type AdminsService struct {
	repo         repository.Admins
	hasher       hash.PasswordHasher
	tokenManager auth.TokenManager
	emailService Emails

	accessTokenTTL  time.Duration
	refreshTokenTTL time.Duration
}

func NewAdminsService(repo repository.Admins, hasher hash.PasswordHasher, tokenManager auth.TokenManager,
	emailService Emails, accessTTL time.Duration, refreshTTL time.Duration) *AdminsService {
	return &AdminsService{
		repo:            repo,
		hasher:          hasher,
		tokenManager:    tokenManager,
		emailService:    emailService,
		accessTokenTTL:  accessTTL,
		refreshTokenTTL: refreshTTL,
	}
}

func (s *AdminsService) SignUp(ctx context.Context, input AdminSignUpInput) error {
	verificationCode := primitive.NewObjectID()
	admin := models.User{
		Name:         input.Name,
		Email:        input.Email,
		Password:     s.hasher.Hash(input.Password),
		DomainId:     primitive.ObjectID{},
		DomainName:   "platform",
		IsAdmin:      true,
		IsBlocked:    false,
		RegisteredAt: time.Now(),
		LastVisitAt:  time.Time{},
		Verification: models.Verification{
			Code: verificationCode,
		},
	}

	if err := s.repo.Create(ctx, admin); err != nil {
		return err
	}

	//return s.emailService.AddToListAdmin(AddToListInput{
	//	Email:            input.Email,
	//	Name:             input.Name,
	//	VerificationCode: verificationCode.Hex(),
	//})

	return nil
}

func (s *AdminsService) GetAll(ctx context.Context) ([]models.User, error) {
	return s.repo.GetAll(ctx)
}

func (s *AdminsService) ChangeAdminBlockStatus(ctx context.Context, editorId string, state string) error {
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

	return s.repo.ChangeBlockStatus(ctx, editorId, stateBool)
}

func (s *AdminsService) ChangeAdminVerifyStatus(ctx context.Context, editorId string, state string) error {
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

	return s.repo.ChangeVerificationStatus(ctx, editorId, stateBool)
}

type UpdateAdminInput struct {
	Id       string
	Name     string
	Email    string
	Password string
	DomainId string
	Verify   bool
	Block    bool
}

func (s *AdminsService) UpdateAdmin(ctx context.Context, newUser UpdateAdminInput) error {
	userId, err := primitive.ObjectIDFromHex(newUser.Id)
	if err != nil {
		return err
	}

	domainId, err := primitive.ObjectIDFromHex(newUser.DomainId)
	if err != nil {
		return err
	}

	oldUser, err := s.repo.GetAdminById(ctx, userId)
	if err != nil {
		return err
	}

	oldUser.Name = newUser.Name
	oldUser.Email = newUser.Email

	if newUser.Password != "" {
		oldUser.Password = s.hasher.Hash(newUser.Password)
	}

	oldUser.DomainId = domainId
	oldUser.Verification.Verified = newUser.Verify
	oldUser.IsBlocked = newUser.Block

	return s.repo.UpdateAdmin(ctx, oldUser)
}

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

// EditorsService - Структура сервиса редакторов
type EditorsService struct {
	repo         repository.Editors
	hasher       hash.PasswordHasher
	tokenManager auth.TokenManager
	emailService Emails

	accessTokenTTL  time.Duration
	refreshTokenTTL time.Duration
}

// NewEditorsService - Создание нового сервиса редакторов
func NewEditorsService(repo repository.Editors, hasher hash.PasswordHasher, tokenManager auth.TokenManager,
	emailService Emails, accessTTL time.Duration, refreshTTL time.Duration) *EditorsService {
	return &EditorsService{
		repo:            repo,
		hasher:          hasher,
		tokenManager:    tokenManager,
		emailService:    emailService,
		accessTokenTTL:  accessTTL,
		refreshTokenTTL: refreshTTL,
	}
}

// TODO uncomment later
//s.emailService.AddToList(AddToListInput{
//Email:            input.Email,
//Name:             input.Name,
//VerificationCode: verificationCode.Hex(),
//})

func (s *EditorsService) SignUp(ctx context.Context, input EditorSignUpInput) error {
	verificationCode := primitive.NewObjectID()
	editor := models.User{
		Name:         input.Name,
		Email:        input.Email,
		Password:     s.hasher.Hash(input.Password),
		DomainId:     input.DomainId,
		DomainName:   input.DomainName,
		IsAdmin:      false,
		IsBlocked:    input.Block,
		RegisteredAt: time.Now(),
		LastVisitAt:  time.Time{},
		Verification: models.Verification{
			Code:     verificationCode,
			Verified: input.Verify,
		},
	}

	if err := s.repo.Create(ctx, editor); err != nil {
		return err
	}

	//return s.emailService.AddToList(AddToListInput{
	//	Email:            editor.Email,
	//	Name:             editor.Name,
	//	VerificationCode: verificationCode.Hex(),
	//})

	return nil
}

//func (s *EditorsService) SignUp(ctx context.Context, input EditorSignUpInput) error {
//	verificationCode := primitive.NewObjectID()
//	editor := models.Editor{
//		Name:         input.Name,
//		Password:     s.hasher.Hash(input.Password),
//		Email:        input.Email,
//		RegisteredAt: time.Now(),
//		LastVisitAt:  time.Now(),
//		UniversityID: input.UniversityID,
//		Verification: models.Verification{
//			Code: verificationCode,
//		},
//	}
//
//	if err := s.repo.Create(ctx, editor); err != nil {
//		return err
//	}
//
//	return
//}

func (s *EditorsService) ChangeEditorBlockStatus(ctx context.Context, editorId string, state string) error {
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

func (s *EditorsService) ChangeEditorVerifyStatus(ctx context.Context, editorId string, state string) error {
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

func (s *EditorsService) GetAllEditors(ctx context.Context) ([]models.User, error) {
	return s.repo.GetAllEditors(ctx)
}

type UpdateEditorInput struct {
	Id         string
	Name       string
	Email      string
	Password   string
	DomainName string
	DomainId   string
	Verify     bool
	Block      bool
}

func (s *EditorsService) UpdateEditor(ctx context.Context, newUser UpdateEditorInput) error {
	userId, err := primitive.ObjectIDFromHex(newUser.Id)
	if err != nil {
		return err
	}

	domainId, err := primitive.ObjectIDFromHex(newUser.DomainId)
	if err != nil {
		return err
	}

	oldUser, err := s.repo.GetEditorById(ctx, userId)
	if err != nil {
		return err
	}

	oldUser.Name = newUser.Name
	oldUser.Email = newUser.Email
	oldUser.Password = s.hasher.Hash(newUser.Password)
	oldUser.DomainId = domainId
	oldUser.Verification.Verified = newUser.Verify
	oldUser.IsBlocked = newUser.Block

	err = s.repo.UpdateEditor(ctx, oldUser)

	return err
}

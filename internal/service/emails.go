package service

import "sas/pkg/email"

type EmailService struct {
	provider email.Provider
	listId   string
}

func NewEmailsService(provider email.Provider, listId string) *EmailService {
	return &EmailService{
		provider: provider,
		listId:   listId,
	}
}

func (s *EmailService) AddToList(input AddToListInput) error {
	return s.provider.AddEmailToList(email.AddEmailInput{
		Email:            input.Email,
		Name:             input.Name,
		VerificationCode: input.VerificationCode,
	})
}

func (s *EmailService) AddToListAdmin(input AddToListInput) error {
	return s.provider.AddEmailToListAdmin(email.AddEmailInput{
		Email:            input.Email,
		Name:             input.Name,
		VerificationCode: input.VerificationCode,
	})
}

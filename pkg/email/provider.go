package email

// AddEmailInput - Добавить поле переменных для кастомизируемых сообщений
type AddEmailInput struct {
	Email            string
	Name             string
	VerificationCode string
}

type Provider interface {
	AddEmailToList(input AddEmailInput) error
}

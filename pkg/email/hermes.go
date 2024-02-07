package email

import (
	"github.com/matcornic/hermes/v2"
)

// TODO Переработать логику составления сообщений
// generateInfo TODO стилизует сообщение в зависимости от того с какого хоста пришло сообщение

func generateInfo() hermes.Hermes {
	return hermes.Hermes{
		Product: hermes.Product{
			Name: "Университетская платформа",
			Link: "http://localhost:8080/ping",
		},
	}
}

func GenerateVerificationEmail(input AddEmailInput) (string, error) {
	h := generateInfo()

	msg := hermes.Email{
		Body: hermes.Body{
			Name: input.Name,
			Intros: []string{
				"Здравствуйте! Вас приветствует команда платформы Университеты!",
			},
			Dictionary: []hermes.Entry{
				{Key: "Name", Value: input.Name},
			},
			Actions: []hermes.Action{
				{
					Instructions: "Для того, чтобы подтвердить регистрацию на платформе необходимо нажать на кнопку ниже!" +
						"или перейти по ссылке внизу сообщения!",
					Button: hermes.Button{
						Text: "Подтвердить регистрацию",
						Link: "http://localhost:8080/api/v1/editors/verify/" + input.VerificationCode,
					},
				},
			},
			Outros: []string{
				"Ссылка для подтверждения перехода, если не работает кнопка:",
				"http://localhost:8080/api/v1/editors/verify/" + input.VerificationCode,
				"Если есть вопросы или вы столкнулись с проблемами пишите на данный адрес:",
				"durdma840@gmail.com",
			},
		},
	}

	//err := os.WriteFile("../test_email.html", []byte(emailBody), 0666)
	//if err != nil {
	//	logger.Errorf("error while writing email %s\n", err)
	//}
	return h.GenerateHTML(msg)
}

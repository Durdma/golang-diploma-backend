package email

import (
	"gopkg.in/gomail.v2"
	"sas/pkg/logger"
	"strconv"
)

type Client struct {
	from     string
	password string
	serv     string
	port     string
}

func NewClient(from string, password string, serv string, port string) *Client {
	return &Client{
		from:     from,
		password: password,
		serv:     serv,
		port:     port,
	}
}

func (c *Client) AddEmailToList(input AddEmailInput) error {
	msg, err := GenerateVerificationEmail(input)
	if err != nil {
		logger.Errorf("error while constructing email %s\n", err)
		return err
	}
	logger.Info("We are here!")
	//return smtp.SendMail(c.serv+":"+c.port,
	//	smtp.PlainAuth("", c.from, c.password, c.serv),
	//	c.from, []string{input.Email}, []byte(mime+msg))

	email := gomail.NewMessage()
	email.SetHeader("From", c.from)
	email.SetHeader("To", input.Email)
	email.SetHeader("Subject", "Verification")
	email.SetBody("text/html", msg)

	port, _ := strconv.Atoi(c.port)

	serv := gomail.NewDialer(c.serv, port, c.from, c.password)

	return serv.DialAndSend(email)
}

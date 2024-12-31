package exchangesmtp

import (
	"net/smtp"
)

// MailSender uses for send plain text emails.
type MailSender struct {
	auth   smtp.Auth
	server string
}

// NewMailSender is constructor for MailSender.
func NewMailSender(username, password, server string) *MailSender {
	return &MailSender{LoginAuth(username, password), server}
}

func (ms *MailSender) sendPlain(mail Mail) error {
	return nil
}

func (ms *MailSender) sendHTML(mail Mail) error {
	return nil
}

// Send is send simple plain text email to list recipients.
func (ms *MailSender) Send(mail Mail) error {
	b, err := mail.ToBytes()
	if err != nil {
		return err
	}

	if err := smtp.SendMail(ms.server, ms.auth, mail.From, mail.To, b); err != nil {
		return err
	}

	return nil
}

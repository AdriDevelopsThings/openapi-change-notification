package email

import (
	"fmt"
	"strings"

	"github.com/adridevelopsthings/openapi-change-notification/config"
	"gopkg.in/mail.v2"
)

type MailDialer struct {
	Dialer *mail.Dialer
	From   string
}

type MailBody struct {
	ContentType string
	Body        string
}

func SendMail(dialer *MailDialer, to string, subject string, body MailBody) error {
	m := mail.NewMessage()
	m.SetHeader("From", dialer.From)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	c := config.GetConfig()
	m.SetHeader("List-Unsubscribe", fmt.Sprintf("<%s>", c.FrontendURL+strings.ReplaceAll(c.FrontendUnsubscribePath, "EMAIL", to)))
	m.SetBody(body.ContentType, body.Body)
	if err := dialer.Dialer.DialAndSend(m); err != nil {
		fmt.Printf("Error while sending E-Mail: %v\n", err)
		return err
	}
	return nil
}

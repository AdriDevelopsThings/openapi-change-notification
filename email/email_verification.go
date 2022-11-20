package email

import (
	"fmt"
	"strings"

	"github.com/adridevelopsthings/openapi-change-notification/config"
)

func SendEmailVerification(d *MailDialer, email string, code string) error {
	return SendMail(d, email, "Verify your email address", MailBody{
		ContentType: "text/plain",
		Body: fmt.Sprintf("Dear user,\n\nplease verify your email address at %s: Just click here: %s",
			config.GetConfig().FrontendURL,
			config.GetConfig().FrontendURL+strings.ReplaceAll(config.GetConfig().FrontendEmailVerificationPath, "CODE", code),
		),
	})
}

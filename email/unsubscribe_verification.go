package email

import (
	"strings"

	"github.com/adridevelopsthings/openapi-change-notification/config"
)

func SendUnsubscribeVerification(d *MailDialer, email string, code string) error {
	return SendMail(d, email, "Verify your unsubscription", MailBody{
		ContentType: "text/plain",
		Body:        "Dear user,\n\nwe are sorry that you want to leave us. Click the following link to finally unsubscribe from your subscriptions: " + config.GetConfig().FrontendURL + strings.ReplaceAll(config.GetConfig().FrontendUnsubscribeVerificationPath, "CODE", code),
	})
}

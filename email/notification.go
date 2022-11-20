package email

import (
	"fmt"
	"strings"

	"github.com/adridevelopsthings/openapi-change-notification/config"
	"github.com/adridevelopsthings/openapi-change-notification/openapi"
)

func SendDeprecationNotification(d *MailDialer, to string, apiMeaning *openapi.OpenAPIMeaning, pathMeaning *openapi.PathMeaning) error {
	c := config.GetConfig()
	return SendMail(d, to, "OpenAPI deprecation notification", MailBody{
		ContentType: "text/plain",
		Body: fmt.Sprintf("Dear user,\n\nwe want to send you a deprecation notification: The path %s with method %s of the served open api configuration on %s was deprecated.\n\nIf you don't want to get emails from us to openapi deprecation changes unsubscribe by clicking here: %s",
			pathMeaning.Path, pathMeaning.Method, apiMeaning.URL,
			c.FrontendURL+strings.ReplaceAll(c.FrontendUnsubscribePath, "EMAIL", to),
		),
	})
}

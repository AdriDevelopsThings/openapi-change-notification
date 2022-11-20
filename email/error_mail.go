package email

import (
	"fmt"

	"github.com/adridevelopsthings/openapi-change-notification/openapi"
)

func SendErroMail(d *MailDialer, to string, apiMeaing *openapi.OpenAPIMeaning, pathMeaning *openapi.PathMeaning) error {
	return SendMail(d, to, "Error while fetching deprecations", MailBody{
		ContentType: "text/plain",
		Body: fmt.Sprintf("Dear user,\n\nwe checked for deprecation updates on served openapi configuration %s on path %s with method %s for 24 hours and got too much errors. We disabled your subscription. Please recheck your configuration and resubscribe or contact our support to solve this issue.",
			apiMeaing.URL,
			pathMeaning.Path,
			pathMeaning.Method,
		),
	})
}

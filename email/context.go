package email

import "context"

const (
	CONTEXT_MAIL_DIALER_NAME = "mail"
)

func GetDialerFromContext(ctx context.Context) *MailDialer {
	return ctx.Value(CONTEXT_MAIL_DIALER_NAME).(*MailDialer)
}

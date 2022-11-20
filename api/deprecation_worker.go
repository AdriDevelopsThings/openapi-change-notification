package api

import (
	"context"
	"fmt"
	"time"

	"github.com/adridevelopsthings/openapi-change-notification/database"
	"github.com/adridevelopsthings/openapi-change-notification/email"
	"github.com/adridevelopsthings/openapi-change-notification/openapi"
)

func StartDeprecationWorker(ctx context.Context) {
	for {
		subscriptions := database.GetFetchingSubscriptions(database.GetDatabaseByContext(ctx))
		for _, subscription := range subscriptions {
			apiMeaning, pathMeaning := subscription.GetMeanings()
			deprecated, err := openapi.GetDeprecated(ctx, apiMeaning, pathMeaning)

			if err != nil {
				fmt.Printf("Error while fetching deprecated for %d: %v\n", subscription.ID, err)
				if subscription.ErrorSince == nil {
					now := time.Now()
					subscription.ErrorSince = &now
				} else if time.Now().Sub(*subscription.ErrorSince).Hours() >= 24 {
					fmt.Printf("Too many errors on %d\n", subscription.ID)
					err := email.SendErroMail(email.GetDialerFromContext(ctx), subscription.MailAddress, apiMeaning, pathMeaning)
					if err != nil {
						fmt.Printf("Error while sending error mail to %s\n", subscription.MailAddress)
					}
					database.GetDatabaseByContext(ctx).Delete(&subscription)
					continue
				}
			} else {
				subscription.ErrorSince = nil
			}

			if deprecated {
				fmt.Printf("Api path of subscription %d is now deprecated.\n", subscription.ID)
				err := email.SendDeprecationNotification(email.GetDialerFromContext(ctx), subscription.MailAddress, apiMeaning, pathMeaning)
				if err != nil {
					fmt.Printf("Error while sending subscription deprecated notification to %s\n", subscription.MailAddress)
					continue
				}
				subscription.Deprecated = true
			}
			subscription.LastFetch = time.Now()
			database.GetDatabaseByContext(ctx).Save(&subscription)
		}
		time.Sleep(time.Minute * 1)
	}
}

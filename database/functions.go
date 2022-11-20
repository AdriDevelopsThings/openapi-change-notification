package database

import (
	"errors"
	"math/rand"
	"time"

	"github.com/adridevelopsthings/openapi-change-notification/email"
	"github.com/adridevelopsthings/openapi-change-notification/openapi"
	"gorm.io/gorm"
)

func Subscribe(db *gorm.DB, email string, apiMeaning *openapi.OpenAPIMeaning, pathMeaning *openapi.PathMeaning, deprecated bool) *Subscription {
	var subscription Subscription
	if err := db.
		Where("email = ?", email).
		Where("openapi_url = ?", apiMeaning.URL).
		Where("path = ?", pathMeaning.Path).
		Where("method = ?", pathMeaning.Method).
		First(&subscription).Error; err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		subscription = Subscription{
			MailAddress: email,
			OpenApiUrl:  apiMeaning.URL,
			Path:        pathMeaning.Path,
			Method:      pathMeaning.Method,
			Deprecated:  deprecated,
		}
		db.Create(&subscription)

	}
	return &subscription
}

func GetSubscriptions(db *gorm.DB, email string) []Subscription {
	var subscriptions []Subscription
	db.Where("email = ?", email).Find(&subscriptions)
	return subscriptions
}

func generateEmailVerificationCode() string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	out := make([]rune, 64)
	rand.Seed(time.Now().UnixNano())
	for i := range out {
		out[i] = letters[rand.Intn(len(letters))]
	}
	return string(out)
}

func VerifiyEmail(db *gorm.DB, mailDialer *email.MailDialer, emailAddress string) *EmailVerification {
	var verification EmailVerification
	if err := db.
		Where("email = ?", emailAddress).
		First(&verification).Error; err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		verification = EmailVerification{
			Code:        generateEmailVerificationCode(),
			Verified:    false,
			MailAddress: emailAddress,
		}
		db.Create(&verification)
		email.SendEmailVerification(mailDialer, emailAddress, verification.Code)
	}
	if time.Now().Sub(verification.CreatedAt).Hours() > 24 {
		db.Delete(&verification)
		return VerifiyEmail(db, mailDialer, emailAddress)
	}
	return &verification
}

func FinishEmailVerification(db *gorm.DB, code string) bool {
	var verification EmailVerification
	if err := db.Where("code = ?", code).
		Where("verified = FALSE").
		First(&verification).Error; err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}
	if time.Now().Sub(verification.CreatedAt).Hours() > 24 {
		db.Delete(&verification)
		return false
	}
	verification.Verified = true
	db.Save(&verification)
	return true
}

func GetFetchingSubscriptions(db *gorm.DB) []Subscription {
	var subscriptions []Subscription
	db.
		Where("subscriptions.last_fetch < ?", time.Now().Add(-time.Hour)).
		Where("subscriptions.deprecated = FALSE").
		Joins("join email_verifications on email_verifications.email = subscriptions.email").
		Where("email_verifications.verified = TRUE").
		Find(&subscriptions)
	return subscriptions
}

func Unsubscribe(db *gorm.DB, mailDialer *email.MailDialer, emailAddress string) *UnsubscribeVerification {
	var verification UnsubscribeVerification
	if err := db.
		Where("email = ?", emailAddress).
		First(&verification).Error; err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		verification = UnsubscribeVerification{
			Code:        generateEmailVerificationCode(),
			MailAddress: emailAddress,
		}
		db.Create(&verification)
		email.SendUnsubscribeVerification(mailDialer, emailAddress, verification.Code)
	}
	if time.Now().Sub(verification.CreatedAt).Hours() > 24 {
		db.Delete(&verification)
		return Unsubscribe(db, mailDialer, emailAddress)
	}
	return &verification
}

func FinishUnsubscribeVerification(db *gorm.DB, code string) bool {
	var verification UnsubscribeVerification
	if err := db.Where("code = ?", code).
		First(&verification).Error; err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}
	if time.Now().Sub(verification.CreatedAt).Hours() > 24 {
		db.Delete(&verification)
		return false
	}
	db.Delete(&verification)
	db.Where("email = ?", verification.MailAddress).Delete(&Subscription{})
	db.Where("email = ?", verification.MailAddress).Delete(&EmailVerification{})
	return true
}

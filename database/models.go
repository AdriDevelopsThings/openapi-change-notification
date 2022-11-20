package database

import (
	"time"

	"github.com/adridevelopsthings/openapi-change-notification/openapi"
)

type EmailVerification struct {
	ID          uint      `gorm:"column:id;primaryKey" json:"id"`
	Code        string    `gorm:"column:code" json:"code"`
	Verified    bool      `gorm:"column:verified" json:"verified"`
	MailAddress string    `gorm:"column:email;unique" json:"email"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
}

type UnsubscribeVerification struct {
	ID          uint      `gorm:"column:id;primaryKey" json:"id"`
	Code        string    `gorm:"column:code" json:"code"`
	MailAddress string    `gorm:"column:email;unique" json:"email"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
}

type Subscription struct {
	ID          uint       `gorm:"column:id;primaryKey" json:"id"`
	MailAddress string     `gorm:"column:email" json:"email"`
	OpenApiUrl  string     `gorm:"column:openapi_url" json:"openapi_url"`
	Path        string     `gorm:"column:path" json:"path"`
	Method      string     `gorm:"column:method" json:"method"`
	Deprecated  bool       `gorm:"column:deprecated" json:"-"`
	LastFetch   time.Time  `gorm:"column:last_fetch;autoCreateTime" json:"-"`
	ErrorSince  *time.Time `gorm:"column:error_since" json:"-"`
	CreatedAt   time.Time  `gorm:"column:created_at;autoCreateTime" json:"created_at"`
}

func (subscription *Subscription) GetMeanings() (*openapi.OpenAPIMeaning, *openapi.PathMeaning) {
	return &openapi.OpenAPIMeaning{URL: subscription.OpenApiUrl},
		&openapi.PathMeaning{Path: subscription.Path, Method: subscription.Method}
}

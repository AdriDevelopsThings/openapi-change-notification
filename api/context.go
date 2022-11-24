package api

import (
	"context"
	"fmt"

	"github.com/adridevelopsthings/openapi-change-notification/config"
	"github.com/adridevelopsthings/openapi-change-notification/database"
	"github.com/adridevelopsthings/openapi-change-notification/email"
	"github.com/adridevelopsthings/openapi-change-notification/openapi"
	"github.com/glebarez/sqlite"
	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
	"gopkg.in/mail.v2"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var redisCache *cache.Cache
var databaseObj *gorm.DB
var mailDialer *email.MailDialer

func BuildNewContext() {
	c := config.GetConfig()
	client := redis.NewClient(&redis.Options{
		Addr: c.RedisURL,
	})
	redisCache = cache.New(&cache.Options{
		Redis: client,
	})

	db, err := gorm.Open(sqlite.Open(c.SQLitePath), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		fmt.Printf("Error while connecting to database: %v\n", err)
		panic("Database error")
	}
	databaseObj = db
	db.AutoMigrate(&database.Subscription{}, &database.EmailVerification{}, &database.UnsubscribeVerification{})
	mailDialer = &email.MailDialer{
		From:   c.SMTPFromAddress,
		Dialer: mail.NewDialer(c.SMTPServer, c.SMTPPort, c.SMTPUsername, c.SMTPPassword),
	}
}

func BuildContext() context.Context {
	if redisCache == nil {
		BuildNewContext()
	}
	ctx := context.TODO()
	ctx = context.WithValue(ctx, openapi.CONTEXT_CACHE_NAME, redisCache)
	ctx = context.WithValue(ctx, database.CONTEXT_DATABASE_NAME, databaseObj)
	ctx = context.WithValue(ctx, email.CONTEXT_MAIL_DIALER_NAME, mailDialer)
	return ctx
}

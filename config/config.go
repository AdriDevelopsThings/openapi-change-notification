package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Config struct {
	Environment                         string
	RedisURL                            string
	SQLitePath                          string
	FrontendStaticServe                 string
	FrontendURL                         string
	FrontendEmailVerificationPath       string
	FrontendUnsubscribePath             string
	FrontendUnsubscribeVerificationPath string
	SMTPServer                          string
	SMTPPort                            int
	SMTPUsername                        string
	SMTPPassword                        string
	SMTPFromAddress                     string
	HCaptchaSecret                      string
}

var CurrentConfig *Config

func LoadConfig() {
	if _, err := os.Stat(".env"); err == nil {
		err := godotenv.Load()
		if err != nil {
			fmt.Printf("Error while reading dotenv: %v\n", err)
			panic("Dotenv reading error")
		}
	}
	CurrentConfig = &Config{
		Environment:                         os.Getenv("ENVIRONMENT"),
		RedisURL:                            os.Getenv("REDIS_URL"),
		SQLitePath:                          os.Getenv("SQLITE_PATH"),
		FrontendStaticServe:                 os.Getenv("FRONTEND_STATIC_SERVE"),
		FrontendURL:                         os.Getenv("FRONTEND_URL"),
		FrontendEmailVerificationPath:       os.Getenv("FRONTEND_EMAIL_VERIFICATION_PATH"),
		FrontendUnsubscribePath:             os.Getenv("FRONTEND_UNSUBSCRIBE_PATH"),
		FrontendUnsubscribeVerificationPath: os.Getenv("FRONTEND_UNSUBSCRIBE_VERIFICATION_PATH"),
		SMTPServer:                          os.Getenv("SMTP_SERVER"),
		SMTPUsername:                        os.Getenv("SMTP_USERNAME"),
		SMTPPassword:                        os.Getenv("SMTP_PASSWORD"),
		SMTPFromAddress:                     os.Getenv("SMTP_FROM_ADDRESS"),
		HCaptchaSecret:                      os.Getenv("HCAPTCHA_SECRET"),
	}
	if CurrentConfig.Environment == "" {
		fmt.Printf("No environment set: Using development\n")
		CurrentConfig.Environment = "development"
	} else if CurrentConfig.Environment != "development" && CurrentConfig.Environment != "production" {
		panic("Set environment to development or production. Your value is wrong.")
	}

	if CurrentConfig.RedisURL == "" {
		fmt.Printf("No redis url set: Setting to localhost:6379\n")
		CurrentConfig.RedisURL = "localhost:6379"
	}

	if CurrentConfig.SQLitePath == "" {
		fmt.Printf("No sqlite path set: Setting to database.db\n")
		CurrentConfig.SQLitePath = "database.db"
	}

	if CurrentConfig.FrontendStaticServe == "" {
		CurrentConfig.FrontendStaticServe = "static"
	}

	if CurrentConfig.FrontendURL == "" {
		fmt.Printf("No frontend URL set: Using development default http://localhost:3000\n")
		CurrentConfig.FrontendURL = "http://localhost:3000"
	}

	if CurrentConfig.FrontendEmailVerificationPath == "" {
		fmt.Printf("No frontend email verification path set: Using development default /email/verify?code=CODE\n")
		CurrentConfig.FrontendEmailVerificationPath = "/email/verify?code=CODE"
	}

	if CurrentConfig.FrontendUnsubscribePath == "" {
		fmt.Printf("No frontend unsubscribe path set: Using development default /unsubscribe?email=EMAIL\n")
		CurrentConfig.FrontendUnsubscribePath = "/unsubscribe?email=EMAIL"
	}

	if CurrentConfig.FrontendUnsubscribeVerificationPath == "" {
		fmt.Printf("No frontend unsubscribe verification path set: Using development default /unsubscribe/verify?code=CODE\n")
		CurrentConfig.FrontendUnsubscribeVerificationPath = "/unsubscribe/verify?code=CODE"
	}

	if CurrentConfig.HCaptchaSecret == "" {
		fmt.Printf("WARNING: No hcaptcha secret set: Hcaptcha verification is not available.\n")
	}

	smtpPort := os.Getenv("SMTP_PORT")
	if smtpPort == "" {
		fmt.Printf("No SMTP port supplied: Using 465 as default\n")
		smtpPort = "465"
	}
	smtpPortParsed, err := strconv.Atoi(smtpPort)
	if err != nil {
		panic("Error while parsing SMTP Port")
	}
	CurrentConfig.SMTPPort = smtpPortParsed

	if CurrentConfig.Environment == "development" {
		fmt.Printf("WARNING: You are running your app in development mode. You should use production mode if you are using this in a production system.\n")
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
}

func GetConfig() *Config {
	if CurrentConfig == nil {
		LoadConfig()
	}
	return CurrentConfig
}

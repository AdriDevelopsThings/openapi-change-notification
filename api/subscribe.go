package api

import (
	"github.com/adridevelopsthings/openapi-change-notification/apierrors"
	"github.com/adridevelopsthings/openapi-change-notification/database"
	"github.com/adridevelopsthings/openapi-change-notification/email"
	"github.com/adridevelopsthings/openapi-change-notification/openapi"
	"github.com/gin-gonic/gin"
)

type SubscribeArguments struct {
	Email      string `json:"email"`
	OpenApiUrl string `json:"openapi_url"`
	Path       string `json:"path"`
	Method     string `json:"method"`
}

func subscribePath(c *gin.Context) {
	var arguments SubscribeArguments
	c.BindJSON(&arguments)
	if !emailRegex.MatchString(arguments.Email) ||
		!openApiUrlRegex.MatchString(arguments.OpenApiUrl) ||
		!pathRegex.MatchString(arguments.Path) ||
		!methodRegex.MatchString(arguments.Method) {
		apierrors.BadRequestError.Abort(c)
		return
	}
	apiMeaning := openapi.OpenAPIMeaning{URL: arguments.OpenApiUrl}
	pathMeaning := openapi.PathMeaning{Path: arguments.Path, Method: arguments.Method}
	ctx := BuildContext()
	deprecated, err := openapi.GetDeprecated(ctx, &apiMeaning, &pathMeaning)
	if err != nil {
		apierrors.AbortError(c, err)
		return
	}
	db := database.GetDatabaseByContext(ctx)
	subscription := database.Subscribe(
		db,
		arguments.Email,
		&apiMeaning,
		&pathMeaning,
		deprecated,
	)
	verified := database.VerifiyEmail(db, email.GetDialerFromContext(ctx), arguments.Email)
	c.JSON(200, gin.H{
		"subscription":   subscription,
		"email_verified": verified.Verified,
	})
}

type UnsubscribeArguments struct {
	SubscribeArguments
	ID string `json:"id"`
}

func unsubscribe(c *gin.Context) {
	emailAddress := c.Query("email")
	if !emailRegex.MatchString(emailAddress) {
		apierrors.BadRequestError.Abort(c)
		return
	}
	ctx := BuildContext()
	_, err := database.Unsubscribe(
		database.GetDatabaseByContext(ctx),
		email.GetDialerFromContext(ctx),
		emailAddress,
	)
	if err != nil {
		err.Abort(c)
		return
	}
	c.JSON(200, gin.H{
		"status": "success",
	})
}

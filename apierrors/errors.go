package apierrors

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type OpenApiChangeNotificationError struct {
	Message    string
	HTTPStatus int
}

func (err *OpenApiChangeNotificationError) Error() string {
	return err.Message
}

func RawAbortError(c *gin.Context, statusCode int, statusMessage string) {
	c.AbortWithStatusJSON(statusCode, gin.H{"error": statusMessage})
}

func (err *OpenApiChangeNotificationError) Abort(c *gin.Context) {
	RawAbortError(c, err.HTTPStatus, err.Message)
}

func AbortError(c *gin.Context, err error) {
	serr, ok := err.(*OpenApiChangeNotificationError)
	if ok {
		serr.Abort(c)
	} else {
		InternalServerError.Abort(c)
		fmt.Printf("Internal server error while handling request. %v\n", err)
	}
}

var InternalServerError = &OpenApiChangeNotificationError{"internal_server_error", 500}
var BadRequestError = &OpenApiChangeNotificationError{"bad_request_error", 400}

var OpenApiFetchingError = &OpenApiChangeNotificationError{"open_api_fetching_error", 400}
var PathCouldNotBeFound = &OpenApiChangeNotificationError{"path_could_not_be_found_error", 400}
var PathMethodCouldNotBeFound = &OpenApiChangeNotificationError{"path_method_could_not_be_found_error", 400}

var SubscriptionNotFound = &OpenApiChangeNotificationError{"subscription_not_found_error", 404}

var EmailVerificationCodeError = &OpenApiChangeNotificationError{"email_verification_code_error", 404}
var UnsubscribeVerificationCodeError = &OpenApiChangeNotificationError{"unsubscribe_verification_code_error", 404}

package api

import (
	"github.com/adridevelopsthings/openapi-change-notification/apierrors"
	"github.com/adridevelopsthings/openapi-change-notification/database"
	"github.com/gin-gonic/gin"
)

func finishEmailVerification(c *gin.Context) {
	code := c.Param("code")
	if code == "" {
		apierrors.BadRequestError.Abort(c)
		return
	}
	ctx := BuildContext()
	success := database.FinishEmailVerification(database.GetDatabaseByContext(ctx), code)
	if success {
		c.JSON(200, gin.H{"status": "success"})
	} else {
		apierrors.EmailVerificationCodeError.Abort(c)
	}
}

func finishUnsubscribeVerification(c *gin.Context) {
	code := c.Param("code")
	if code == "" {
		apierrors.BadRequestError.Abort(c)
		return
	}
	ctx := BuildContext()
	success := database.FinishUnsubscribeVerification(database.GetDatabaseByContext(ctx), code)
	if success {
		c.JSON(200, gin.H{"status": "success"})
	} else {
		apierrors.UnsubscribeVerificationCodeError.Abort(c)
	}
}

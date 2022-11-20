package database

import (
	"context"

	"gorm.io/gorm"
)

const (
	CONTEXT_DATABASE_NAME = "database"
)

func GetDatabaseByContext(ctx context.Context) *gorm.DB {
	return ctx.Value(CONTEXT_DATABASE_NAME).(*gorm.DB)
}

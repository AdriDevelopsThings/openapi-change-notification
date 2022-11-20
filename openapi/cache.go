package openapi

import (
	"context"

	"github.com/go-redis/cache/v8"
)

const (
	CONTEXT_CACHE_NAME = "cache"
)

func getCacheByContext(ctx context.Context) *cache.Cache {
	return ctx.Value(CONTEXT_CACHE_NAME).(*cache.Cache)
}

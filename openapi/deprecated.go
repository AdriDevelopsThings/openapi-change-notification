package openapi

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/adridevelopsthings/openapi-change-notification/apierrors"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-redis/cache/v8"
)

type PathMeaning struct {
	Path   string `json:"path"`
	Method string `json:"method"`
}

func (meaning *PathMeaning) RedisKey(apiMeaing *OpenAPIMeaning) string {
	return "deprecated_" + apiMeaing.URL + meaning.Method + meaning.Path
}

func (meaning *PathMeaning) Translate(openapi *openapi3.T) {
	for _, server := range openapi.Servers {
		if len(server.URL) > 10 && strings.HasPrefix(meaning.Path, server.URL) {
			meaning.Path = meaning.Path[len(server.URL):]
		}
	}
}

func GetDeprecated(ctx context.Context, apiMeaning *OpenAPIMeaning, pathMeaning *PathMeaning) (bool, error) {
	redis_cache := getCacheByContext(ctx)

	openapi, err := GetOpenApi(ctx, apiMeaning)
	if err != nil {
		return false, err
	}
	pathMeaning.Translate(openapi)

	var deprecated bool
	redis_key := pathMeaning.RedisKey(apiMeaning)
	err = redis_cache.Get(ctx, redis_key, &deprecated)
	if err == cache.ErrCacheMiss {
		pathMeaning.Translate(openapi)
		item := openapi.Paths.Find(pathMeaning.Path)
		if item == nil {
			return false, apierrors.PathCouldNotBeFound
		}
		method := item.Operations()[strings.ToUpper(pathMeaning.Method)]
		if method == nil {
			return false, apierrors.PathMethodCouldNotBeFound
		}
		deprecated = method.Deprecated
		if err = redis_cache.Set(&cache.Item{
			Ctx:   ctx,
			Key:   redis_key,
			Value: deprecated,
			TTL:   time.Hour,
		}); err != nil {
			fmt.Printf(
				"Error while setting openapi deprecated redis cache for url %s: path %s method %s: %v\n",
				apiMeaning.URL,
				pathMeaning.Path,
				pathMeaning.Method,
				err,
			)
		}

	} else if err != nil {
		return false, err
	}
	return deprecated, nil
}

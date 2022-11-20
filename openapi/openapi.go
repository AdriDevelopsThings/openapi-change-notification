package openapi

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/adridevelopsthings/openapi-change-notification/apierrors"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-redis/cache/v8"
)

type OpenAPIMeaning struct {
	URL string `json:"url"`
}

func (meaning *OpenAPIMeaning) RedisKey() string {
	return "openapi_" + meaning.URL
}

func parseOpenApi(content []byte) (*openapi3.T, error) {
	return openapi3.NewLoader().LoadFromData(content)
}

func GetOpenApi(ctx context.Context, meaning *OpenAPIMeaning) (*openapi3.T, error) {
	redis_cache := getCacheByContext(ctx)
	var openapi []byte
	err := redis_cache.Get(ctx, meaning.RedisKey(), &openapi)
	if err == nil {
		return parseOpenApi(openapi)
	} else if err == cache.ErrCacheMiss {
		resp, err := http.Get(meaning.URL)
		if err != nil {
			return nil, err
		}
		if resp.StatusCode != 200 {
			return nil, apierrors.OpenApiFetchingError
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		if err = redis_cache.Set(&cache.Item{
			Ctx:   ctx,
			Key:   meaning.RedisKey(),
			Value: body,
			TTL:   time.Hour,
		}); err != nil {
			fmt.Printf("Error while setting openapi response redis cache for url %s: %v\n", meaning.URL, err)
		}
		return parseOpenApi(body)
	} else {
		return nil, err
	}
}

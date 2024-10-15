package resourcesv2service

import (
	"context"
	"encoding/json"
	"github.com/NorskHelsenett/ror/cmd/api/apiconnections"
	"time"

	"github.com/NorskHelsenett/ror/pkg/helpers/kvcachehelper"
	"github.com/NorskHelsenett/ror/pkg/helpers/kvcachehelper/rediscache"
	"github.com/NorskHelsenett/ror/pkg/rlog"
	"github.com/NorskHelsenett/ror/pkg/rorresources"
)

var ResourceCache *resourceCache

type resourceCache struct {
	cacheLayers []kvcachehelper.CacheInterface
}

func GetResourceCache() *resourceCache {
	if ResourceCache == nil {
		ResourceCache = &resourceCache{
			cacheLayers: []kvcachehelper.CacheInterface{
				rediscache.NewRedisCache(apiconnections.RedisDB, kvcachehelper.CacheOptions{Prefix: "res_uid_", Timeout: time.Hour * 1}),
			},
		}
	}
	return ResourceCache
}

func (rc *resourceCache) Set(ctx context.Context, resource *rorresources.Resource) {
	key := resource.GetUID()
	value, err := json.Marshal(resource)
	if err != nil {
		rlog.Error("Could not marshal resource", err, rlog.Any("resource", resource.GetAPIVersion()), rlog.Any("kind", resource.GetKind()), rlog.Any("name", resource.GetName()), rlog.Any("error", err))
	}
	for _, c := range rc.cacheLayers {
		c.Set(ctx, key, string(value))
	}
}

func (rc *resourceCache) Get(ctx context.Context, key string) *rorresources.Resource {
	for _, c := range rc.cacheLayers {
		data, cacheHit := c.Get(ctx, key)
		if cacheHit {
			resource := &rorresources.Resource{}
			err := json.Unmarshal([]byte(data), resource)
			if err != nil {
				rlog.Error("Could not unmarshal resource", err, rlog.String("uid", key), rlog.Any("error", err))
			}
			return rorresources.NewResourceFromStruct(*resource)
		}
	}
	return nil
}

func (rc *resourceCache) Remove(ctx context.Context, key string) {
	for _, c := range rc.cacheLayers {
		c.Remove(ctx, key)
	}
}

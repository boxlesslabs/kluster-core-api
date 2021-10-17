package goCache

import (
	"encoding/json"
	"errors"
	"github.com/patrickmn/go-cache"
	"time"
)

type (
	AppCache interface {
		Set(key string, data interface{}, expiration time.Duration) error
		Get(key string) ([]byte, error)
	}

	appCache struct {
		client *cache.Cache
	}
)

func NewAppCache() *appCache {
	return &appCache{cache.New(24*time.Hour, 12*time.Hour)}
}

func (cache *appCache) Set(key string, data interface{}, expiration time.Duration) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if err := cache.client.Add(key, b, expiration); err != nil {
		return err
	}
	return nil
}

func (cache *appCache) Get(key string) ([]byte, error) {
	res, exist := cache.client.Get(key)
	if !exist {
		return nil, nil
	}

	resByte, ok := res.([]byte)
	if !ok {
		return nil, errors.New("Format is not arr of bytes")
	}

	return resByte, nil
}
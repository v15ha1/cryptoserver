package hitbtc

import (
	"encoding/json"
	"errors"
	"time"
	"github.com/patrickmn/go-cache"
)

var MyCache CacheItf

type CacheItf interface {
	Set(key string, data interface{} ) error
	Get(key string) ([]byte, error)
}

type AppCache struct {
	client *cache.Cache
}

func (r *AppCache) Set(key string, data interface{}) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	
	r.client.Set(key, b, 60*time.Minute)
	return nil
}

func (r *AppCache) Get(key string) ([]byte, error) {
	res, exist := r.client.Get(key)
	if !exist {
		return nil, nil
	}

	resByte, ok := res.([]byte)
	if !ok {
		return nil, errors.New("Format is not arr of bytes")
	}

	return resByte, nil
}

func InitCache() {
	MyCache = &AppCache{
		client: cache.New(5*time.Minute, 10*time.Minute),
	}
}
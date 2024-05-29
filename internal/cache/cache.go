package cache

import (
	"time"

	"github.com/patrickmn/go-cache"
)

type Cache struct {
	cache *cache.Cache
}

func NewCache(defaultExpiration, cleanupInterval time.Duration) *Cache {
	c := cache.New(defaultExpiration, cleanupInterval)
	return &Cache{cache: c}
}

func (c *Cache) Get(key string) (string, bool) {
	value, found := c.cache.Get(key)
	if found {
		return value.(string), true
	}
	return "", false
}

func (c *Cache) Set(key string, value string, expiration time.Duration) {
	c.cache.Set(key, value, expiration)
}

func (c *Cache) Delete(key string) {
	c.cache.Delete(key)
}

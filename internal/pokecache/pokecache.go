package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
	mu    sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

const (
	logCache bool = false
)

func NewCache(interval time.Duration) *Cache {
	cache := Cache{
		cache: map[string]cacheEntry{},
		mu:    sync.Mutex{},
	}
	ticker := time.NewTicker(interval)
	go func() {
		for range ticker.C {
			cache.reapLoop(interval)
		}
	}()
	return &cache
}

func (c *Cache) Get(key string) ([]byte, bool) {
	value, ok := c.get(key)
	if logCache {
		CacheLog(ok, key)
	}
	return value, ok
}

func (c *Cache) get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	//try read cache
	var entry []byte
	if entry, ok := c.cache[key]; ok {
		return entry.val, true
	}
	return entry, false
}

func (c *Cache) Add(key string, value []byte) error {
	//read the cache to check for existing key
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, ok := c.cache[key]; ok {
		return fmt.Errorf("Cache value for key: %s already exists", key)
	}

	//lock the cache for a write
	//add the key or return error
	//unlock the cache

	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val:       value,
	}

	return nil
}

func (c *Cache) reapLoop(interval time.Duration) {
	//clean up expired cache entries
	c.mu.Lock()
	defer c.mu.Unlock()
	for key, item := range c.cache {
		expireTime := item.createdAt.Add(interval)
		if time.Now().After(expireTime) {
			delete(c.cache, key)
		}
	}
}

func CacheLog(hit bool, key string) {
	if hit {
		fmt.Printf("Cache hit on key '%s'\n", key)
	} else {
		fmt.Printf("Cache miss on key '%s'\n", key)
	}
}

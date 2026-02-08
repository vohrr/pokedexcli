package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
	mu    *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) {

}

func (c Cache) Get(key string) ([]byte, bool) {

	return []byte{}, false
}

func (c Cache) Add(key string, value []byte) error {
	//read the cache to check for existing key
	//lock the cache for a write
	//add the key or return error
	//unlock the cache

	return nil
}

func (c Cache) reapLoop() {

}

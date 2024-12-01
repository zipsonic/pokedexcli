package pokecache

import "time"

func NewCache(interval time.Duration) *Cache {

	var cache Cache

	cache.entries = make(map[string]cacheEntry)

	go cache.reapLoop(interval)

	return &cache
}

package pokecache

import "time"

func NewCache(interval time.Duration) Cache {

	var PCache Cache

	go PCache.reapLoop(interval)

	return PCache
}

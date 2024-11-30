package pokecache

import "time"

var cache Cache

func NewCache(interval time.Duration) {

	go cache.reapLoop(interval)

	return
}

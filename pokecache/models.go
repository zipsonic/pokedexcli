package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	mu      sync.Mutex
	entries map[string]cacheEntry
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry := c.entries[key]

	entry.createdAt = time.Now()
	entry.val = val

	c.entries[key] = entry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	returnVal := make([]byte, 0)
	entry, ok := c.entries[key]
	if ok {
		returnVal = entry.val
	}

	return returnVal, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	for {
		c.mu.Lock()

		for key := range c.entries {
			entry := c.entries[key]
			if entry.createdAt.Add(interval).Before(time.Now()) {
				delete(c.entries, key)
			}
		}
		c.mu.Unlock()
		time.Sleep(interval)
	}
}

package pokecache

import (
	"time"
	"sync"
)

type Cache struct {
	entries map[string]cacheEntry
	mu sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val []byte
}

func NewCache(interval time.Duration) *Cache {
	var cache Cache
	cache.entries = make(map[string]cacheEntry)
	go cache.reapLoop(interval)
	return &cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	var entry cacheEntry
	entry.createdAt = time.Now()
	entry.val = val

	c.entries[key] = entry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	
	entry, ok := c.entries[key]
	if ok {
		return entry.val, true
	} else {
		return nil, false
	}
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	for range ticker.C {
		c.mu.Lock()
		for str, entry := range c.entries {
			if time.Since(entry.createdAt) > interval {
				delete(c.entries, str)
			}
		}
		c.mu.Unlock()
	}

}
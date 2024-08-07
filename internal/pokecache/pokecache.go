package pokecache

import (
	"time"
)

const DefaultTime = time.Millisecond * 10

type Cache struct {
	cache map[string]cacheEntry
	//mu    *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval ...time.Duration) Cache {
	var actualInterval time.Duration
	if len(interval) > 0 {
		actualInterval = interval[0]
	} else {
		actualInterval = DefaultTime
	}
	c := Cache{
		cache: make(map[string]cacheEntry),
	}
	go c.reapLoop(actualInterval)
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.cache[key] = cacheEntry{
		val:       val,
		createdAt: time.Now(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	cacheEntry, ok := c.cache[key]
	return cacheEntry.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(interval)
	}
}

func (c *Cache) reap(interval time.Duration) {
	cutoffTime := time.Now().Add(-interval)
	for k, v := range c.cache {
		if v.createdAt.Before(cutoffTime) {
			delete(c.cache, k)
		}
	}
}

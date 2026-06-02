package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache    map[string]cacheEntry
	interval time.Duration
	mu       sync.RWMutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = cacheEntry{
		val:       val,
		createdAt: time.Now(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	_, ok := c.cache[key]
	if !ok {
		return nil, false
	}
	return c.cache[key].val, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	for range ticker.C {
		c.mu.Lock()

		for key, val := range c.cache {
			duration := time.Since(val.createdAt)
			if c.interval < duration {
				delete(c.cache, key)
			}
		}

		c.mu.Unlock()
	}
}

func NewCache(inter time.Duration) *Cache {
	c := Cache{
		cache:    make(map[string]cacheEntry),
		interval: inter,
	}

	go c.reapLoop()
	return &c
}

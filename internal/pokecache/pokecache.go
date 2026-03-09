package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	mu      sync.Mutex
	entries map[string]cacheEntry
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func (e cacheEntry) isExpired(expirationTime time.Duration) bool {
	return time.Since(e.createdAt) > expirationTime
}

func NewCache(expirationTime time.Duration) *Cache {
	newCache := &Cache{
		entries: map[string]cacheEntry{},
	}
	go newCache.reapLoop(expirationTime)
	return newCache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) (val []byte, ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.entries[key]
	return entry.val, ok
}

func (c *Cache) reapLoop(expirationTime time.Duration) {
	ticker := time.NewTicker(expirationTime)
	for range ticker.C {
		c.mu.Lock()
		for key, entry := range c.entries {
			if entry.isExpired(expirationTime) {
				delete(c.entries, key)
			}
		}
		c.mu.Unlock()
	}
}

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
	cachedValues map[string]cacheEntry
	mutex        *sync.Mutex
}

func NewCache(interval time.Duration) Cache {
	mux := &sync.Mutex{}
	cache := Cache{}
	cache.mutex = mux
	cache.cachedValues = make(map[string]cacheEntry)
	go cache.reapLoop(interval)
	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.cachedValues[key] = cacheEntry{time.Now(), val}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	cachedVal, exists := c.cachedValues[key]
	return cachedVal.val, exists
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now(), interval)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	for k, v := range c.cachedValues {
		if v.createdAt.Before(now.Add(-last)) {
			delete(c.cachedValues, k)
		}
	}
}

package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
    cache map[string]cacheEntry
    mu    *sync.RWMutex
}

type cacheEntry struct {
    createdAt   time.Time 
    val         []byte
}


func NewCache(interval time.Duration) Cache {
    c := Cache {
        cache:  make(map[string]cacheEntry),
        mu:     &sync.RWMutex{}, 
    }
    go c.reapLoop(interval)
    return c
}

func (c *Cache) Add(key string, val []byte) {
    c.mu.Lock()
    defer c.mu.Unlock()

    entry := cacheEntry{
        createdAt:  time.Now(),
        val:        val,
    }
    c.cache[key] = entry
}

func (c *Cache) Get(key string) ([]byte, bool) {
    c.mu.RLock()
    defer c.mu.RUnlock()
    entry, ok := c.cache[key]
    return entry.val, ok 
}

func (c *Cache) reapLoop(interval time.Duration) {
    ticker := time.NewTicker(interval)
    defer ticker.Stop()
    timeNow := time.Now()
    for range ticker.C {
        c.reap(&timeNow)
    }
}

func (c *Cache) reap(prevTime *time.Time){
    c.mu.Lock()
    defer c.mu.Unlock()
    for k, v := range c.cache {
        if v.createdAt.Before(*prevTime) {
            delete(c.cache, k)
        }
    }
    *prevTime = time.Now() 
}
